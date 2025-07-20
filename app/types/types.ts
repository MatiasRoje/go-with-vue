export type User = {
  email: string
  first_name: string
  last_name: string
}

export type BackendResponse<T> = {
  error: boolean
  message: string
  data?: T
}

export type BackendErrorResponse = {
  error: boolean
  message: string
}