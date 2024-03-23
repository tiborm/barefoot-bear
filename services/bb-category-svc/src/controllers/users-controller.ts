import { Request, Response, NextFunction } from 'express'
import usersSvc from '../services/users-service'

const getUsers = (req: Request, res: Response, next: NextFunction) => {
  res.json(usersSvc.getUsers())
}
const getUsersByID = (req: Request, res: Response, next: NextFunction) => {
  res.json(usersSvc.getUserByID(req.params.id))
}


export default {
  getUsers,
  getUsersByID,
}
