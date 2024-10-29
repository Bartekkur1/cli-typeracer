import { describe, expect, test } from 'bun:test';
import { ServerClient } from './util';
import { Command } from './types';

describe('Welcome Message test', () => {

  test('Should receive welcome message', async () => {
    const client = await ServerClient.createPlayer();
    await client.register();

    const message = await client.getLatestMessage();
    expect(message.command).toBe(Command.Welcome);
    expect(message.content).toBe('Welcome to the server!');
    expect(message.playerId).toBeString();

    await client.close();
  })
});