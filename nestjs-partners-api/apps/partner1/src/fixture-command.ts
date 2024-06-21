import { NestFactory } from '@nestjs/core';
import { Partner1Module } from './partner1.module';
import { fixture } from '@app/core/fixture';

async function bootstrap() {
  const app = await NestFactory.createApplicationContext(Partner1Module);

  const events = [
    {
      id: '10853e59-dc5b-4d7b-a028-01513ef50d76',
      name: 'Event 001 - Partner1',
      description: 'Event 001 Description - Partner1',
      price: 100,
      date: '2021-10-10T10:00:00',
    },
    {
      id: 'e0352b32-7698-4805-b029-28302b3a911f',
      name: 'Event 002 - Partner1',
      description: 'Event 002 Description - Partner1',
      price: 200,
      date: '2021-10-10T12:00:00',
    },
  ];

  await fixture(app, events, 10);

  console.log('Partner1 - data created');
  await app.close();
}
bootstrap();
