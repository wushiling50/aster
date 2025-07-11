import re

import httpx
from bs4 import BeautifulSoup

import timezone_util  # 自定义时区工具库
import pycountry  # 国家信息库

from email import guess_by_profile_email  # 导入邮箱推测模块

def guess_by_profile_timezone(handle: str):
    url = f"https://github.com/{handle}"
    html_content = httpx.get(url).text
    soup = BeautifulSoup(html_content, 'html.parser')

    profile_timezone = soup.find('profile-timezone')
    if profile_timezone is None:
        return None
    
    timezone = profile_timezone.text
    m = re.findall(r'([+-]\d{2}:\d{2})', timezone) # +08:00
    if m is None or len(m) != 1:
        return None

    timezone: str = m[0].replace(':', '') # +0800
    return timezone_util.get_countries_by_timezone(timezone)

def guess_by_profile_country_name(handle: str):
    url = f"https://github.com/{handle}"
    html_content = httpx.get(url).text
    soup = BeautifulSoup(html_content, 'html.parser')

    profile_name = soup.find('span', class_='p-name vcard-fullname d-block overflow-hidden')
    profile_bio = soup.find('div', class_='p-note user-profile-bio mb-3 js-user-profile-bio f4')
    profile_readme = soup.find('turbo-frame', class_='user-profile-frame')

    if profile_name is None or profile_bio is None or profile_readme is None:
        return None
    
    text = profile_name.text + profile_bio.text + profile_readme.text

    countries = find_country(text)

    return countries

def find_country(text):
    full_list =[]
    countries = sorted(pycountry.countries, key=lambda x: -len(x))
    for country in countries:
        if country.alpha_2.lower() in text.lower():
            full_list.append(country)
    return full_list