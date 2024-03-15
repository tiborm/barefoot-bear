import User from '../models/user-model'

export default class UsersService {

  static getUsers() {
    const users = new Set<User>([
      { id: 'uniqueIDSample-1', name: 'Morgan', age: 24 },
      { id: 'uniqueIDSample-2', name: 'Sarah', age: 54 },
      { id: 'uniqueIDSample-3', name: 'John', age: 43 },
      { id: 'uniqueIDSample-4', name: 'Megan', age: 98 },
    ])

    return [...users]
  }

}
