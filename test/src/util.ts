import { Command, type Message } from "./types";

export const spawnWS = () => {
  return new WebSocket('ws://localhost:8080');
}

export const createMessage = ({
  playerId,
  command,
  content,
}: Message): string => {
  return JSON.stringify({
    playerId,
    command,
    content,
  });
};

export const wait = async (ms: number) => {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve(undefined);
    }, ms);
  });
};

export class ServerClient {
  private ws: WebSocket;

  playerId: string;
  gameId?: string;
  messages: Message[];

  constructor() {
    this.ws = spawnWS();
    this.playerId = "not_empty";
    this.messages = [];
  }

  async startAndWait() {
    return new Promise((resolve) => {
      this.ws.addEventListener('message', (event) => {
        this.messages.push(JSON.parse(event.data as string) as Message);
      });
      this.ws.addEventListener('open', () => {
        resolve(this.ws as WebSocket);
      });
    });
  }

  static async createPlayer(params?: { register?: boolean }): Promise<ServerClient> {
    const client = new ServerClient();
    await client.startAndWait();
    if (params?.register) {
      await client.register();
    }
    return client;
  }

  async sendMessage(message: Omit<Message, 'playerId'>): Promise<void> {
    return new Promise(async (resolve) => {
      this.ws.send(createMessage({
        playerId: this.playerId,
        ...message
      }));
      await wait(100);
      resolve(undefined);
    });
  }

  async getLatestMessage(): Promise<Message> {
    return new Promise(async (resolve) => {
      if (this.messages.length === 0) {
        this.ws.addEventListener('message', (event) => {
          resolve(JSON.parse(event.data as string) as Message);
        });
      } else {
        await wait(100);
        resolve(this.messages[this.messages.length - 1]);
      }
    });
  }

  async close() {
    return new Promise((resolve) => {
      this.ws.close();
      resolve(undefined);
    });
  }

  async register() {
    await this.sendMessage({
      command: Command.Welcome,
      content: ``,
    });
    const response = await this.getLatestMessage();
    this.playerId = response.playerId;
  }

  async createGame(): Promise<string> {
    await this.sendMessage({
      command: Command.CreateGame,
      content: ``,
    });
    const response = await this.getLatestMessage();
    if (response.command === Command.GameCreated) {
      this.gameId = response.content;
    }
    return response.content;
  }

  async joinGame(gameId: string): Promise<void> {
    await this.sendMessage({
      command: Command.JoinGame,
      content: gameId,
    });
    const response = await this.getLatestMessage();
    if (response.command === Command.GameJoined) {
      this.gameId = gameId;
    }
  }

  async readyCheck(ready: boolean = true): Promise<void> {
    await this.sendMessage({
      command: ready ? Command.PlayerReady : Command.PlayerNotReady,
      content: ``,
    });
  }

  async startGame(): Promise<void> {
    await this.sendMessage({
      command: Command.StartGame,
      content: ''
    });
  }

  async leaveGame(): Promise<void> {
    await this.sendMessage({
      command: Command.PlayerLeave,
      content: ''
    });
  }
}