import { describe, expect, test } from 'bun:test';
import { createPlayer, webSocket } from './util';
import { Command } from './types';
import { v4 } from 'uuid';

describe('Create game', () => {

  test('Should not allow game creation without welcoming', async () => {
    const ws = webSocket();
    await ws.startAndWait();
    const response = await ws.sendMessage({
      command: Command.CreateGame,
      content: ``,
    });

    expect(response.command).toBe(Command.Error);
    expect(response.content).toBe("player not found");

    await ws.close();
  });

  test('Should create game and return its invite token', async () => {
    const player = await createPlayer();
    const createGameResponse = await player.sendMessage({
      command: Command.CreateGame,
      content: ``,
    });

    expect(createGameResponse.command).toBe(Command.GameCreated);
    expect(createGameResponse.content).not.toBeEmpty();
    expect(createGameResponse.content.length).toEqual(5);

    await player.close();
  });

});