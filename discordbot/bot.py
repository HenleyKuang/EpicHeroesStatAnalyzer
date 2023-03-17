# bot.py
import os

import discord
from discord.ext import commands

TOKEN = os.environ['DISCORD_TOKEN']

client = discord.Client(intents=discord.Intents.default())
client = commands.Bot(intents=discord.Intents.default(), command_prefix="!*")


@client.event
async def on_ready():
    print(f'{client.user} has connected to Discord!')

@client.command()
async def eval(ctx):
  await ctx.send('Test!')

client.run(TOKEN)