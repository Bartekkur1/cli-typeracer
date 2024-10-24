import { describe, expect, test } from 'bun:test';
import { webSocket } from './util';
import { Command } from './types';

describe('Welcome Message test', () => {

  test('Should receive welcome message', async (done) => {
    const ws = webSocket();
    await ws.startAndWait();
    const response = await ws.sendMessage({
      playerId: "",
      command: Command.Welcome,
      content: `${Date.now()}`,
    });

    expect(response.playerId).toBeString();
    expect(response.command).toBe(Command.Welcome);
    expect(response.content).toBe('Welcome to the server!');

    await ws.close();
    done();
  })
});