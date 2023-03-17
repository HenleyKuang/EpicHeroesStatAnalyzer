# bot.py
import os

import discord
from discord.ext import commands

TOKEN = os.environ["DISCORD_TOKEN"]

intents = discord.Intents.default()

client = discord.Client(intents=intents)
client = commands.Bot(intents=intents, command_prefix="!*")


@client.event
async def on_ready():
    print(f"{client.user} has connected to Discord!")


@client.command()
async def eval(ctx):
    await ctx.send("Test!")


@client.event
async def on_message(message):
    # https://discordpy.readthedocs.io/en/stable/api.html?highlight=on_message#discord.Message
    username = str(message.author).split("#")[0]
    channel = str(message.channel.name)
    user_message = str(message.content)
    attachments = message.attachments
    # Attachments: [<Attachment id=1086357302585593990 filename='toko_stats.jpg' url='https://cdn.discordapp.com/attachments/1086348096872661022/1086357302585593990/toko_stats.jpg'>]
    mentions = message.mentions
    # Mentions: [<Member id=1086148392092172358 name='Prometheus Bot' discriminator='7031' bot=True nick=None guild=<Guild id=1047094083518218272 name='Prometheus & Deucalion & Aidos union server' shard_id=0 chunked=False member_count=74>>]

    if username != "YouSnoozeYouLose":
        return

    toko_response = {
        "Estimated Dmg": {
            "Basic Atk DMG": 1051512,
            "Basic Atk DMG with Crit": 2167271,
            "Passive Atk Dmg": 1051512,
            "Passive Atk Dmg with Crit": 2167271,
            "Skill Atk Dmg": 7003070,
            "Skill Atk Dmg with Crit": 14434028,
        },
        "Hero": "Samurai Girl",
        "Stats": {
            "ATK": 637280,
            "Accuracy": 46,
            "Armor": 4468,
            "Block": 5.6,
            "Broken Armor": 65,
            "Crit": 81,
            "Crit DMG": 31,
            "Crit Damage Resistance": 15,
            "Crit Resistance": 4,
            "DMG Immune": 23,
            "Dodge": 0,
            "Effect Hit": 24,
            "Effect Res": 28,
            "HP": 11285601,
            "Hit": 0,
            "Holy DMG": 0,
            "Power": 6278276,
            "Skill DMG": 33,
            "Speed": 133,
        },
    }

    print(f"Message {user_message} by {username} on {channel}")
    print(f"Attachments: {attachments}")
    print(f"Mentions: {mentions}")

    if message.author == client.user:
        return

    if channel == "bot-testing":
        await message.channel.send(toko_response)


client.run(TOKEN)
