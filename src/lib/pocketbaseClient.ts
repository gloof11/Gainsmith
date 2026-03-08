import PocketBase from 'pocketbase'

let temppb =  new PocketBase(`${import.meta.env.VITE_APP_POCKETBASE_URL}`)
temppb.autoCancellation(false)

export const pb = temppb
