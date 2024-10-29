export enum Command {
  Error = "ERROR",
  Unrecognize = "UNRECOGNIZED",
  Welcome = "WELCOME",
  CreateGame = "CREATE_GAME",
  GameCreated = "GAME_CREATED",
  JoinGame = "JOIN_GAME",
  GameJoined = "GAME_JOINED",
  ACK = 'ACK',
  StartGame = "START_GAME",
  GameStarting = "GAME_STARTING",
  GameStarted = "GAME_STARTED",
  PlayerJoined = "PLAYER_JOINED",
  PlayerReady = "PLAYER_READY",
  PlayerNotReady = "PLAYER_NOT_READY",
  PlayerLeft = "PLAYER_LEFT",
  PlayerLeave = "PLAYER_LEAVE",
  GameClosed = "GAME_CLOSED",
}

export interface Message {
  playerId: string;
  command: Command;
  content: string;
}