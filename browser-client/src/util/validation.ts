
const username= '^[a-zA-Z0-9_]{5,16}$'
const password='^.{8,16}$'


export function getRegex(key:string) {
    switch (key) {
        case 'username':
            return username
        case 'password':
            return password
        default:
            return null
    }
}
