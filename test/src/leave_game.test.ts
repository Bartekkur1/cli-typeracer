import { describe, expect, test } from 'bun:test';
import { Command } from './types';
import { ServerClient } from './util';

describe('Leave Game', () => {

  test('Should notify game owner about opponent leaving', async () => {
    const host = await ServerClient.createPlayer({ register: true });
    const opponent = await ServerClient.createPlayer({ register: true });

    await host.createGame();
    await opponent.joinGame(host.gameId!);

    await opponent.leaveGame();

    const opponentLeftMessage = host.messages.find(m => m.command === Command.PlayerLeft);
    expect(opponentLeftMessage).toBeDefined();
    expect(opponentLeftMessage!.content).toEqual(opponent.playerId!);
  });

  test('Should notify opponent about game owner leaving', async () => {
    const host = await ServerClient.createPlayer({ register: true });
    const opponent = await ServerClient.createPlayer({ register: true });

    await host.createGame();
    await opponent.joinGame(host.gameId!);

    await host.leaveGame();

    const hostLeftMessage = host.messages.find(m => m.command === Command.GameClosed);
    expect(hostLeftMessage).toBeDefined();
    expect(hostLeftMessage!.content).toEqual(host.playerId!);
  });

});