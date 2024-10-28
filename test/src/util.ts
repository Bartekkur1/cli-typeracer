import { Command, type Message } from "./types";

export const spawnWS = () => {
  return new WebSocket('ws://localhost:8080');
}

export const createMessage = ({
  playerId,
  command,
  content,
}: Message): string => {
  const mess = JSON.stringify({
    playerId,
    command,
    content,
  });
  // console.log(mess);
  return mess;
};

export const wait = async (ms: number) => {
  return new Promise((resolve) => {
    setTimeout(() => {
      resolve(undefined);
    }, ms);
  });
};

export const createPlayer = async () => {
  const connection = webSocket();
  await connection.startAndWait();
  const welcomeResponse = await connection.sendMessage({
    command: Command.Welcome,
    content: "",
  });
  if (welcomeResponse.command !== Command.Welcome) {
    throw new Error('Player creation failed!');
  }
  connection.setPlayerId(welcomeResponse.playerId);
  return connection;
};

export const webSocket = () => ({
  ws: spawnWS(),
  playerId: "not_empty",
  setPlayerId: function (playerId: string) {
    this.playerId = playerId;
  },
  startAndWait: async function () {
    return new Promise((resolve) => {
      this.ws.addEventListener('open', () => {
        resolve(this.ws);
      });
    });
  },
  sendMessage: async function (message: Omit<Message, "playerId">): Promise<Message> {
    return new Promise((resolve) => {
      this.ws.addEventListener('message', (event) => {
        resolve(JSON.parse(event.data as string) as Message);
      });
      this.ws.send(createMessage({
        playerId: this.playerId,
        ...message,
      }));
    });
  },
  waitForMessage: async function (): Promise<Message> {
    return new Promise((resolve) => {
      this.ws.addEventListener('message', (event) => {
        resolve(JSON.parse(event.data as string) as Message);
      });
    });
  },
  close: async function () {
    return new Promise((resolve) => {
      this.ws.close();
      resolve(undefined);
    });
  }
});