import { describe, expect, test } from 'bun:test';
import { createPlayer } from './util';
import { Command } from './types';

describe('Ready Check', () => {

  test('Should join game and change self state to ready/not ready', async () => {
    const player = await createPlayer();
    const createGameResponse = await player.sendMessage({
      command: Command.CreateGame,
      content: ``,
    });

    expect(createGameResponse.command).toBe(Command.GameCreated);
    expect(createGameResponse.content).not.toBeEmpty();
    expect(createGameResponse.content.length).toEqual(5);

    const readyResponse = await player.sendMessage({
      command: Command.Ready,
      content: ``,
    });

    expect(readyResponse.command).toBe(Command.Error);
    expect(readyResponse.content).toBe(`waiting for opponent`);

    const notReadyResponse = await player.sendMessage({
      command: Command.NotReady,
      content: ``,
    });

    expect(notReadyResponse.command).toBe(Command.Error);
    expect(notReadyResponse.content).toBe(`waiting for opponent`);

    await player.close();
  });

  test('Should not be able to be ready without opponent in game', async () => {
    const player = await createPlayer();
    const createGameResponse = await player.sendMessage({
      command: Command.CreateGame,
      content: ``,
    });

    expect(createGameResponse.command).toBe(Command.GameCreated);
    expect(createGameResponse.content).not.toBeEmpty();
    expect(createGameResponse.content.length).toEqual(5);

    const readyResponse = await player.sendMessage({
      command: Command.Ready,
      content: ``,
    });

    expect(readyResponse.command).toBe(Command.Error);
    expect(readyResponse.content).toBe(`waiting for opponent`);
    await player.close();
  });

  test('Should be able to ready up with opponent in game', async () => {
    const host = await createPlayer();
    const opponent = await createPlayer();

    const createGameResponse = await host.sendMessage({
      command: Command.CreateGame,
      content: ``,
    });

    expect(createGameResponse.command).toBe(Command.GameCreated);
    expect(createGameResponse.content).not.toBeEmpty();
    expect(createGameResponse.content.length).toEqual(5);

    const gameId = createGameResponse.content;

    const joinGameResponse = await opponent.sendMessage({
      command: Command.JoinGame,
      content: gameId,
    });

    expect(joinGameResponse.command).toBe(Command.GameJoined);

    const hostReadyResponse = await host.sendMessage({
      command: Command.Ready,
      content: ``,
    });

    expect(hostReadyResponse.command).toBe(Command.ACK);
    expect(hostReadyResponse.content).toBe(`Player ${host.playerId} is ready`);

    const opponentReadyResponse = await opponent.sendMessage({
      command: Command.Ready,
      content: ``,
    });

    expect(opponentReadyResponse.command).toBe(Command.ACK);
    expect(opponentReadyResponse.content).toBe(`Player ${opponent.playerId} is ready`);

    const opponentNotReadyResponse = await opponent.sendMessage({
      command: Command.NotReady,
      content: ``,
    });

    expect(opponentNotReadyResponse.command).toBe(Command.ACK);
    expect(opponentNotReadyResponse.content).toBe(`Player ${opponent.playerId} is not ready`);

    const hostNotReadyResponse = await opponent.sendMessage({
      command: Command.NotReady,
      content: ``,
    });
    expect(hostNotReadyResponse.command).toBe(Command.ACK);
    expect(hostNotReadyResponse.content).toBe(`Player ${opponent.playerId} is not ready`);

    await host.close();
    await opponent.close();
  });

});