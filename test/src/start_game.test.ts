import { describe, expect, test } from 'bun:test';
import { createPlayer, webSocket } from './util';
import { Command } from './types';
import { v4 } from 'uuid';

describe('Create game', () => {

  test('Should not be able to start game without oponent', async () => {
    const player = await createPlayer();

    const createGameResponse = await player.sendMessage({
      command: Command.CreateGame,
      content: ""
    });

    expect()

    player.sendMessage({
      command: Command.StartGame,
      content: ""
    });
  })

});