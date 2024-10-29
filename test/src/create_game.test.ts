import { describe, expect, test } from 'bun:test';
import { Command } from './types';
import { v4 } from 'uuid';
import { ServerClient } from './util';

describe('Create game', () => {

  test('Should not allow game creation without welcoming', async () => {
    const player = await ServerClient.createPlayer();
    await player.createGame();

    const response = await player.getLatestMessage();
    expect(response.command).toBe(Command.Error);
    expect(response.content).toBe('player not found');
  });

  test('Should create game and return its invite token', async () => {
    const player = await ServerClient.createPlayer();
    await player.register();
    await player.createGame();

    const response = await player.getLatestMessage();
    expect(response.command).toBe(Command.GameCreated);
    expect(response.content).toBeString();
    expect(response.content.length).toEqual(5);
  });

});