import { describe, expect, test } from 'bun:test';
import { Command } from './types';
import { ServerClient, wait } from './util';

describe('Start game', () => {

  test('Should not be able to start game without opponent', async () => {
    const host = await ServerClient.createPlayer({ register: true });
    await host.createGame();
    await host.startGame();

    const latestMessage = await host.getLatestMessage();
    expect(latestMessage.command).toBe(Command.Error);
    expect(latestMessage.content).toBe("waiting for opponent");

    await host.close();
  });

  test('Should not be able to start game without opponent ready', async () => {
    const host = await ServerClient.createPlayer({ register: true });
    const opponent = await ServerClient.createPlayer({ register: true });

    await host.createGame();
    await opponent.joinGame(host.gameId!);

    await host.startGame();
    const latestMessage = await host.getLatestMessage();
    expect(latestMessage.command).toBe(Command.Error);
    expect(latestMessage.content).toBe("players are not ready");

    await host.close();
    await opponent.close();
  });

  test('Players should be notified about game starting', async () => {
    const host = await ServerClient.createPlayer({ register: true });
    const opponent = await ServerClient.createPlayer({ register: true });

    await host.createGame();
    await opponent.joinGame(host.gameId!);

    await host.readyCheck(true);
    await opponent.readyCheck(true);

    await host.startGame();

    expect(host.messages.some(m => m.command === Command.GameStarting)).toBeTrue();
    expect(opponent.messages.some(m => m.command === Command.GameStarting)).toBeTrue();

    await wait(5100);
    expect(host.messages.some(m => m.command === Command.GameStarted)).toBeTrue();
    expect(opponent.messages.some(m => m.command === Command.GameStarted)).toBeTrue();
  });

});