import { describe, expect, test } from 'bun:test';
import { Command } from './types';
import { ServerClient, wait } from './util';

describe('Ready Check', () => {

  test('Should join game and change self state to ready/not ready', async () => {
    const player = await ServerClient.createPlayer({ register: true });
    await player.createGame();
    await player.readyCheck();

    let latestResponse = await player.getLatestMessage();
    expect(latestResponse.command).toBe(Command.Error);
    expect(latestResponse.content).toBe(`waiting for opponent`);

    await player.readyCheck(false);
    latestResponse = await player.getLatestMessage();
    expect(latestResponse.command).toBe(Command.Error);
    expect(latestResponse.content).toBe(`waiting for opponent`);

    await player.close();
  });

  test('Should be able to ready up with opponent in game', async () => {
    const host = await ServerClient.createPlayer({ register: true });
    const opponent = await ServerClient.createPlayer({ register: true });

    await host.createGame();
    await opponent.joinGame(host.gameId!);

    await host.readyCheck();
    await opponent.readyCheck();

    const hostReadyNotification = host.messages.filter(m => m.command === Command.PlayerReady);
    expect(hostReadyNotification.length).toBe(2);

    const opponentReadyNotification = host.messages.filter(m => m.command === Command.PlayerReady);
    expect(opponentReadyNotification.length).toBe(2);

    await host.close();
    await opponent.close();
  });

});