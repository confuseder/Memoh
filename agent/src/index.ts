import { Elysia } from 'elysia'
import { chatModule } from './modules/chat'
import { corsMiddleware } from './middlewares/cors'
import { errorMiddleware } from './middlewares/error'

const app = new Elysia()
  .use(corsMiddleware)
  .use(errorMiddleware)
  .use(chatModule)
  .listen(8081)

console.log(
  `Agent Gateway is running at ${app.server?.hostname}:${app.server?.port}`
)
