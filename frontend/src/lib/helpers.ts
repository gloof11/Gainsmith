export const rules = [
  (value: string) => {
    if (value) return true
    return 'Required'
  }
]
