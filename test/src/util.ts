import { type Message } from "./types";

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

export const webSocket = () => ({
  ws: spawnWS(),
  startAndWait: async function () {
    return new Promise((resolve) => {
      this.ws.addEventListener('open', () => {
        resolve(this.ws);
      });
    });
  },
  sendMessage: async function (message: Message): Promise<Message> {
    return new Promise((resolve) => {
      this.ws.addEventListener('message', (event) => {
        resolve(JSON.parse(event.data as string) as Message);
      });
      this.ws.send(createMessage(message));
    });
  },
  close: async function () {
    return new Promise((resolve) => {
      this.ws.close();
      resolve(undefined);
    });
  }
});