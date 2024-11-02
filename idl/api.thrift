namespace go api

struct User {
    1: required string name,
    2: required double score,
    3: required string nation,
    4: required string confidence, // confidence level
}

struct TalentRankRequest {
    1: required string username,
}

struct TalentRankResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required User user_info,
}

struct SearchRequest {
    1: required string keyword,
    2: optional list<string> domain_list, // Differentiate using programming languages
    3: optional list<string> nation_list,
    4: optional i64 page_num,
}

struct SearchResponse {
    1: required i64 status_code = 0,
    2: optional string status_msg,
    3: required list<User> user_list,
}

service BasicService {
    TalentRankResponse TalentRank(1: TalentRankRequest req) (api.get="/aster/talent/rank/")
    SearchResponse Search(1: SearchRequest req) (api.get="/aster/search")
}