import { describe, expect, test } from 'bun:test';
import { Command } from './types';
import { ServerClient, wait } from './util';

describe('Join Game', () => {

  test('Should not allow game join without welcoming', async () => {
    const host = await ServerClient.createPlayer({ register: true });
    await host.createGame();

    const opponent = await ServerClient.createPlayer();
    await opponent.joinGame(host.gameId!);

    const response = await opponent.getLatestMessage();
    expect(response.command).toBe(Command.Error);
    expect(response.content).toBe('player not found');

    await host.close();
    await opponent.close();
  });

  test('Should be able to join existing game', async () => {
    const host = await ServerClient.createPlayer({ register: true });
    const opponent = await ServerClient.createPlayer({ register: true });

    await host.createGame();
    const createGameResponse = await host.getLatestMessage();

    expect(createGameResponse.command).toBe(Command.GameCreated);
    expect(createGameResponse.content).toBeString();

    await opponent.joinGame(host.gameId!);

    const joinResponse = await opponent.getLatestMessage();
    expect(joinResponse.content).toBe("");
    expect(joinResponse.command).toBe(Command.GameJoined);

    await host.close();
    await opponent.close();
  });

  test('Should notify game owner about player joining', async () => {
    const host = await ServerClient.createPlayer({ register: true });
    const opponent = await ServerClient.createPlayer({ register: true });

    await host.createGame();
    await host.getLatestMessage();

    await opponent.joinGame(host.gameId!);

    const latestHostMessage = await host.getLatestMessage();
    expect(latestHostMessage.command).toBe(Command.PlayerJoined);

    await host.close();
    await opponent.close();
  });

});