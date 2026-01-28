import { generateText, ModelMessage, stepCountIs, streamText, TextStreamPart, ToolSet } from 'ai'
import { createChatGateway } from './gateway'
import { ClientType, Schedule } from './types'
import { system, schedule } from './prompts'

export interface AgentParams {
  apiKey: string
  baseUrl: string
  model: string
  clientType: ClientType
  locale?: Intl.LocalesArgument
  language?: string
  maxSteps?: number
  maxContextLoadTime: number
  platforms?: string[]
  currentPlatform?: string
}

export interface AgentInput {
  messages: ModelMessage[]
  query: string
}

export interface AgentResult {
  messages: ModelMessage[]
}

export const createAgent = (params: AgentParams) => {
  const gateway = createChatGateway(params.clientType)
  const messages: ModelMessage[] = []

  const maxSteps = params.maxSteps ?? 50

  const generateSystem = () => {
    return system({
      date: new Date(),
      locale: params.locale,
      language: params.language,
      maxContextLoadTime: params.maxContextLoadTime,
      platforms: params.platforms ?? [],
      currentPlatform: params.currentPlatform,
    })
  }

  const ask = async (input: AgentInput): Promise<AgentResult> => {
    messages.push(...input.messages)
    messages.push({
      role: 'user',
      content: input.query,
    })
    const { response } = await generateText({
      model: gateway({
        apiKey: params.apiKey,
        baseURL: params.baseUrl,
      })(params.model),
      system: generateSystem(),
      stopWhen: stepCountIs(maxSteps),
      messages,
    })
    return {
      messages: response.messages,
    }
  }

  async function* stream(input: AgentInput): AsyncGenerator<TextStreamPart<ToolSet>, AgentResult> {
    messages.push(...input.messages)
    messages.push({
      role: 'user',
      content: input.query,
    })
    const { response, fullStream } = streamText({
      model: gateway({
        apiKey: params.apiKey,
        baseURL: params.baseUrl,
      })(params.model),
      system: generateSystem(),
      stopWhen: stepCountIs(maxSteps),
      messages,
    })
    for await (const event of fullStream) {
      yield event
    }
    return {
      messages: (await response).messages,
    }
  }

  const triggerSchedule = async (
    input: AgentInput,
    scheduleData: Schedule
  ) => {
    messages.push(...input.messages)
    messages.push({
      role: 'user',
      content: schedule({
        schedule: scheduleData,
        locale: params.locale,
        date: new Date(),
      }),
    })
    return await ask(input)
  }

  return {
    ask,
    stream,
    triggerSchedule,
  }
}