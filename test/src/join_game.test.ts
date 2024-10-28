import { describe, expect, test } from 'bun:test';
import { createPlayer, webSocket } from './util';
import { Command } from './types';

describe('Join Game', () => {

  test('Should not allow game join without welcoming', async () => {
    const ws = webSocket();
    await ws.startAndWait();
    const response = await ws.sendMessage({
      command: Command.JoinGame,
      content: `123123`,
    });

    expect(response.command).toBe(Command.Error);
    expect(response.content).toBe("player not found");

    await ws.close();
  });

  test('Should be able to join existing game', async () => {
    const host = await createPlayer();
    const opponent = await createPlayer();

    const createGameResponse = await host.sendMessage({
      command: Command.CreateGame,
      content: ``
    });

    expect(createGameResponse.command).toBe(Command.GameCreated);
    expect(createGameResponse.content).toBeString();

    const joinResponse = await opponent.sendMessage({
      command: Command.JoinGame,
      content: createGameResponse.content
    });

    expect(joinResponse.command).toBe(Command.GameJoined);

    await host.close();
    await opponent.close();
  });

});