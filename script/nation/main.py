import json
import sys

from github import Github, Auth
from loguru import logger
from ruamel import yaml

from aster_profile import (
    guess_by_profile_timezone, 
    guess_by_profile_country_name
)

from aster_email import guess_by_profile_email

# 不同推测方法的权重分配
EMAIL_GUESS_WEIGHT = float(0.15)
PROFILE_TIMEZONE_GUESS_WEIGHT = float(0.20)
COUNTRY_NAME_GUESS_WEIGHT = float(0.25)

# 验证权重总和为0.6
assert EMAIL_GUESS_WEIGHT + PROFILE_TIMEZONE_GUESS_WEIGHT + COUNTRY_NAME_GUESS_WEIGHT == 0.6

def main():
    config_path = "config/config.yaml"

    with open(config_path,"r") as f:
        config = yaml.load(f,Loader=yaml.RoundTripLoader)

    auth = Auth.Token(config["GithubAPIToken"])

    argv = sys.argv

    if len(argv) != 2:
        logger.error("Usage: Python Argv Error")
        return

    USERNAME = argv[1]

    confidence_dict = dict()
    with Github(auth=auth) as g:
        email_guess = guess_by_profile_email(USERNAME, g)
        if email_guess is not None:
            for country in email_guess:
                confidence_dict[country] = confidence_dict.get(country, 0) + EMAIL_GUESS_WEIGHT

        profile_timezone_guess = guess_by_profile_timezone(USERNAME)
        if profile_timezone_guess is not None:
            for country in profile_timezone_guess:
                confidence_dict[country] = confidence_dict.get(country, 0) + PROFILE_TIMEZONE_GUESS_WEIGHT

        country_name_guess = guess_by_profile_country_name(USERNAME)
        if country_name_guess is not None:
            for country in country_name_guess:
                confidence_dict[country] = confidence_dict.get(country, 0) + COUNTRY_NAME_GUESS_WEIGHT

        print(json.dumps(
                    confidence_dict,
                    separators=(',', ': ')
                ))
        
if __name__ == '__main__':
    main()