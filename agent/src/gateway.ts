import { createGateway as createAiGateway } from 'ai'
import { createOpenAI } from '@ai-sdk/openai'
import { createAnthropic } from '@ai-sdk/anthropic'
import { createGoogleGenerativeAI } from '@ai-sdk/google'
import { ClientType } from './types'

export const createChatGateway = (clientType: ClientType) => {
  const clients = {
    [ClientType.OPENAI]: createOpenAI,
    [ClientType.ANTHROPIC]: createAnthropic,
    [ClientType.GOOGLE]: createGoogleGenerativeAI,
  }
  return (clients[clientType] ?? createAiGateway)
}