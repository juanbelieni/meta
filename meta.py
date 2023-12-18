#!/bin/env python

import yaml
import subprocess
import re
import sys
from yaml.loader import SafeLoader
from pathlib import Path
from dataclasses import dataclass


@dataclass
class Arg:
    name: str


@dataclass
class Script:
    command: str
    args: list[Arg]


@dataclass
class Meta:
    scripts: dict[Script]


def read_config_file() -> dict:
    """
    Read config files from one of the following paths:
    - meta.yaml
    - meta.yml
    """

    paths = [Path("meta.yaml"), Path("meta.yml")]
    path = next((path for path in paths if path.exists()), None)

    if path is None:
        str_paths = ", ".join(map(str, paths))
        raise Exception(
            f"None of the following meta config files found in this directory: {str_paths}"
        )

    with path.open("r") as file:
        config = yaml.load(file, Loader=SafeLoader)

        if config is None:
            raise Exception(
                f"Config file {str(path)} is empty or could not be read"
            )

        return config


def parse_config(config: dict) -> Meta:
    """
    Parse config dict into a Meta instance
    """

    scripts = dict()

    for script_name, script_config in config.items():
        args = []

        if type(script_config) == str:
            # If config is a string, treat it as the command
            script_command = script_config
        elif type(script_config) == dict:
            script_command = script_config["command"]

            if "args" in script_config:
                for arg_name in script_config["args"]:
                    if type(arg_name) == str:
                        args.append(Arg(arg_name))
                    else:
                        raise Exception(
                            f"Could not parse arg {script_name}: {arg_name}"
                        )
        else:
            raise Exception(
                f"Could not parse script {script_name}: {script_config}"
            )

        script = Script(script_command, args)
        scripts[script_name] = script

    return Meta(scripts)


def run_command(meta: Meta, script_name: str, script_arg_values: list[str]):
    """
    Run command by replacing argument templates with the specified values
    """
    match meta.scripts.get(script_name):
        case None:
            raise Exception(f"Unknown script name: {script_name}")
        case script:
            command = script.command
            pattern = r'\{\{(.*?)\}\}'
            matches = re.findall(pattern, command)

            script_arg_dict = dict(zip(
                [arg.name for arg in script.args],
                script_arg_values
            ))

            for arg_name in matches:
                if arg_name not in script_arg_dict:
                    raise Exception(f"Argument not specified: {arg_name}")

                command = command.replace(
                    f"{{{{{arg_name}}}}}",
                    script_arg_dict[arg_name]
                )

            subprocess.call(command, shell=True)


if __name__ == "__main__":
    config = read_config_file()
    meta = parse_config(config)

    if len(sys.argv) == 1:
        raise Exception("Script name not specified")

    script_name = sys.argv[1]
    script_arg_values = sys.argv[2:]

    run_command(meta, script_name, script_arg_values)
