export const rules = [
  (value: any) => {
    if (value) return true
    return 'Required'
  }
]
