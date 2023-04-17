# bot.py
import os
import re

import discord
import json
import requests
from discord.ext import commands

TOKEN = os.environ["DISCORD_TOKEN"]

intents = discord.Intents.default()

client = discord.Client(intents=intents)
client = commands.Bot(intents=intents, command_prefix="!*")

HERO_ANALYSIS_ENDPOINT = "http://127.0.0.1:3000/heroanalysis"
COMMANDS_ENDPOINT = "http://127.0.0.1:3000/commands"

COMMAND_REGEX = r"\<\@[0-9]+\>\s\!(?P<command>[a-z]+)"


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

    is_mentioned = False
    for mention in mentions:
        if mention.name == "Prometheus Bot":
            is_mentioned = True
            break
    if is_mentioned == False:
        return

    print(f"Message {user_message} by {username} on {channel}")
    print(f"Attachments: {attachments}")
    print(f"Mentions: {mentions}")

    p = re.compile(COMMAND_REGEX)
    m = p.search(user_message)
    if m is None:
        response = get_all_commands()
        await message.reply(response)
        return

    command = m.group("command")
    if command:
        print(f"Command: {command}")

    if message.author == client.user:
        return

    if command == "help":
        response = get_all_commands()
        await message.reply(response)
    elif len(attachments) > 0:
        image_url = attachments[0].url
        response = get_image_stats(image_url, command)
        if command == "debug":  # user requesting to debug
            await message.reply(response)
        elif "message" in response:  # has error in response
            await message.reply(response)
        else:
            await message.reply(format_reply(response))


def format_reply(response):
    print(response)
    # response = {
    #     "Estimated DMG": {
    #         "Basic Atk DMG": 1051512,
    #         "Basic Atk DMG with Crit": 2167271,
    #         "Passive Atk DMG": 1051512,
    #         "Passive Atk DMG with Crit": 2167271,
    #         "Skill Atk DMG": 7003070,
    #         "Skill Atk DMG with Crit": 14434028,
    #     },
    #     "Hero": "Samurai Girl",
    #     "Stats": {
    #         "ATK": 637280,
    #         "Accuracy": 46,
    #         "Armor": 4468,
    #         "Block": 5.6,
    #         "Broken Armor": 65,
    #         "Crit": 81,
    #         "Crit DMG": 31,
    #         "Crit Damage Resistance": 15,
    #         "Crit Resistance": 4,
    #         "DMG Immune": 23,
    #         "Dodge": 0,
    #         "Effect Hit": 24,
    #         "Effect Res": 28,
    #         "HP": 11285601,
    #         "Hit": 0,
    #         "Holy DMG": 0,
    #         "Power": 6278276,
    #         "Skill DMG": 33,
    #         "Speed": 133,
    #     },
    # }
    crit_rate = response["Stats"]["Crit"]
    basic_dmg = response["Estimated DMG"]["Basic Atk DMG"]
    basic_dmg_crit = response["Estimated DMG"]["Basic Atk DMG with Crit"]
    passive_dmg = response["Estimated DMG"]["Passive Atk DMG"]
    passive_dmg_crit = response["Estimated DMG"]["Passive Atk DMG with Crit"]
    skill_dmg = response["Estimated DMG"]["Skill Atk DMG"]
    skill_dmg_crit = response["Estimated DMG"]["Skill Atk DMG with Crit"]
    # reply = (
    #     f"Your hero has an estimated basic attack damage of **{basic_dmg:,}**, passive attack damage of **{passive_dmg:,}**, and skill attack damage of **{skill_dmg:,}**. "
    #     + f"With a **{crit_rate}%** chance of CRIT, your hero's basic attack damage would increase to **{basic_dmg_crit:,}**, passive attack damage increases to **{passive_dmg_crit:,}**, "
    #     + f"and skill attack damage increases to **{skill_dmg_crit:,}**.\n"
    #     + "Note: This does not account for several factors including enemy's DMG Immune, your hero's passive buffs, or their skill dmg multipliers!"
    # )
    reply = (
    f"```╔═════════════════════════════════════════════════╗\n"
    +   "║            Prometheus Damage Analysis           ║\n"
    +   "╠═════════════╤═══════════╤═══════════╤═══════════╣\n"
    +   "║ Crit?       │   Basic   │  Passive  │   Skill   ║\n"
    +   "╟─────────────┼───────────┼───────────┼───────────╢\n"
    +  f"║ No Crit     │{basic_dmg:11,d}|{passive_dmg:11,d}|{skill_dmg:11,d}║\n"
    +   "╟─────────────┼───────────┼───────────┼───────────╢\n"
    +  f"║ Crit        |{basic_dmg_crit:11,d}|{passive_dmg_crit:11,d}|{skill_dmg_crit:11,d}║\n"
    +   "╚═════════════╧═══════════╧═══════════╧═══════════╝\n```"
    + "Note: This does not account for several factors including enemy's DMG Immune, your hero's passive buffs, or their skill dmg multipliers!"
    )
    return reply


def get_image_stats(image_url, command):
    data = {
        "imageURL": image_url,
        "heroName": command,  # command may be "!toko" (just "toko" gets extracted as the command string though) which is the hero name.
    }
    response = requests.post(url=HERO_ANALYSIS_ENDPOINT, data=data)
    return json.loads(response.text)


def get_all_commands():
    response = requests.get(url=COMMANDS_ENDPOINT)
    return json.loads(response.text)


client.run(TOKEN)
