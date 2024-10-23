export enum Command {
  Error = "ERROR",
  Unrecognize = "UNRECOGNIZED",
  Welcome = "WELCOME",
  CreateGame = "CREATE_GAME",
  GameCreated = "GAME_CREATED",
  JoinGame = "JOIN_GAME",
  GameJoined = "GAME_JOINED"
}

export interface Message {
  playerId: string;
  command: Command;
  content: string;
}