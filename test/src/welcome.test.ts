import { describe, test } from 'bun:test';
import { webSocket } from './util';
import { v4 } from 'uuid';
import { Command } from './types';

describe('Welcome Message test', () => {

  test('Should receive welcome message', async (done) => {
    const playerId = v4();

    const ws = webSocket();
    await ws.startAndWait();
    const response = await ws.sendMessage({
      playerId,
      command: Command.Welcome,
      content: `${Date.now()}`,
    });

    console.log(response);

    await ws.close();
    done();
  })
});