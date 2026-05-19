import PocketBase from 'pocketbase'

const pb = new PocketBase(
  import.meta.env.VITE_APP_POCKETBASE_URL || window.location.origin,
)
pb.autoCancellation(false)

export default pb
