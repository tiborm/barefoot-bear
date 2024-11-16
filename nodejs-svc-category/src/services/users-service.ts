import User from "../models/user-model";

export default class UsersService {
    static getUsers() {
        const users = new Set<User>([
            { id: "1234", name: "Morgan", age: 24 },
            { id: "1235", name: "Sarah", age: 54 },
            { id: "1236", name: "John", age: 43 },
            { id: "1237", name: "Megan", age: 98 },
        ]);

        return [...users];
    }

    static getUserByID(id: string) {
        return [...UsersService.getUsers()].find((user) => user.id === id);
    }
}
