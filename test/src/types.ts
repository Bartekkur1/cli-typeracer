export enum Command {
  Error = "ERROR",
  Unrecognize = "UNRECOGNIZED",
  Welcome = "WELCOME",
  CreateGame = "CREATE_GAME",
  GameCreated = "GAME_CREATED",
  JoinGame = "JOIN_GAME",
  GameJoined = "GAME_JOINED",
  Ready = "READY",
  NotReady = "NOT_READY",
  ACK = 'ACK',
  StartGame = "START_GAME",
  GameStarted = "GAME_STARTED",
}

export interface Message {
  playerId: string;
  command: Command;
  content: string;
}