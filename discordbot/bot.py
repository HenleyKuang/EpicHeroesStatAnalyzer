# bot.py
import os

import discord

TOKEN = os.environ['DISCORD_TOKEN']

client = discord.Client(intents=discord.Intents.default())

@client.event
async def on_ready():
    print(f'{client.user} has connected to Discord!')

client.run(TOKEN)