import { afterAll, describe, expect, test } from 'bun:test';
import { createPlayer, webSocket } from './util';
import { Command } from './types';
import { v4 } from 'uuid';

describe('Start game', () => {

  test('Should not be able to start game without opponent', async () => {
    const player = await createPlayer();

    const createGameResponse = await player.sendMessage({
      command: Command.CreateGame,
      content: ""
    });

    expect(createGameResponse.command).toBe(Command.GameCreated);

    const startGameResponse = await player.sendMessage({
      command: Command.StartGame,
      content: ""
    });

    expect(startGameResponse.command).toBe(Command.Error);
    expect(startGameResponse.content).toBe("waiting for opponent");

    await player.close();
  });

  test('Should not be able to start game without opponent ready', async () => {
    const host = await createPlayer();
    const opponent = await createPlayer();

    const createGameResponse = await host.sendMessage({
      command: Command.CreateGame,
      content: ""
    });

    expect(createGameResponse.command).toBe(Command.GameCreated);

    const joinGameResponse = await opponent.sendMessage({
      command: Command.JoinGame,
      content: createGameResponse.content
    });

    expect(joinGameResponse.command).toBe(Command.GameJoined);

    const startGameResponse = await host.sendMessage({
      command: Command.StartGame,
      content: ""
    });

    expect(startGameResponse.command).toBe(Command.Error);
    expect(startGameResponse.content).toBe("players are not ready");

    await host.close();
    await opponent.close();
  });

  test('Should be able to start game with opponent ready', async () => {
    const host = await createPlayer();
    const opponent = await createPlayer();

    const createGameResponse = await host.sendMessage({
      command: Command.CreateGame,
      content: ""
    });
    expect(createGameResponse.command).toBe(Command.GameCreated);

    const joinGameResponse = await opponent.sendMessage({
      command: Command.JoinGame,
      content: createGameResponse.content
    });
    expect(joinGameResponse.command).toBe(Command.GameJoined);

    const hostReadyResponse = await host.sendMessage({
      command: Command.Ready,
      content: ""
    });
    expect(hostReadyResponse.command).toBe(Command.ACK);

    const opponentReadyResponse = await opponent.sendMessage({
      command: Command.Ready,
      content: ""
    });
    expect(opponentReadyResponse.command).toBe(Command.ACK);

    host.waitForMessage().then((message) => {
      expect(message.command).toBe(Command.GameStarted);
      expect(message.content.length).toEqual(13);
    });

    await host.sendMessage({
      command: Command.StartGame,
      content: ""
    });

    await host.close();
    await opponent.close();
  });

});