from datetime import datetime

import pytz

# { +0800: ['XXX/XXXXXXXX', ...], ...}
TIMEZONE_OFFSETS = dict()
# { 'America/Los_Angeles': ['us', ...], ...}
TZ_COUNTRIES = dict()


def init_timezone_offsets():
    global TIMEZONE_OFFSETS
    for tz in pytz.all_timezones:
        timezone = pytz.timezone(tz)
        # 获取当前时区的当前时间
        local_time = datetime.now(timezone)
        # 计算UTC偏移量
        utc_offset = local_time.utcoffset()
        assert utc_offset is not None
        # 将偏移量转换为分钟
        minutes, remainder = divmod(utc_offset.total_seconds(), 60)
        assert remainder == 0
        # 将分钟转换为小时和分钟
        hours, minutes = divmod(minutes, 60)
        hours_abs = abs(hours)
        # 确定符号（东半球为+，西半球为-）
        sign = "-" if hours < 0 else "+"
        # 格式化为"±HHMM"形式
        hours_formatted = sign + str(int(hours_abs)).zfill(2) + str(int(minutes)).zfill(2)
        assert len(hours_formatted) == 5
        # 将时区添加到对应偏移量的列表中
        if hours_formatted not in TIMEZONE_OFFSETS:
            TIMEZONE_OFFSETS[hours_formatted] = list()
        TIMEZONE_OFFSETS[hours_formatted].append(tz)


def init_tz_counties():
    global TZ_COUNTRIES
    # 遍历所有国家和对应的时区
    for country, timezones in pytz.country_timezones.items():
        # 遍历该国家的所有时区
        for tz in timezones:
            if tz not in TZ_COUNTRIES:
                TZ_COUNTRIES[tz] = list()
            # 将国家代码（小写）添加到时区的国家列表中
            TZ_COUNTRIES[tz].append(country.lower())


init_timezone_offsets()
init_tz_counties()

def get_countries_by_timezone(tz: str):
    timezones = TIMEZONE_OFFSETS[tz]
    countries = set()
    for timezone in timezones:
        if timezone in TZ_COUNTRIES:
            countries.update(TZ_COUNTRIES[timezone])
    return countries