import { Request, Response, NextFunction } from 'express'
import usersSvc from '../services/users-service'

const getUsers = (req: Request, res: Response, next: NextFunction) => {
  res.json(usersSvc.getUsers())
}

export default {
  getUsers
}
