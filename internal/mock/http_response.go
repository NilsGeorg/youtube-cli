package mock

import "fmt"

var VideoResponse = fmt.Sprintf(`{
    "kind": "youtube#videoListResponse",
    "etag": "22LdEv8Ng8-Ad68aTiG9vLh04eY",
	"nextPageToken": "",
	"prevPageToken": "",
	"regionCode": "",
    "items": %s,
    "pageInfo": {
        "totalResults": 1,
        "resultsPerPage": 1
    }
}`, VideoResponseItems)

var VideoResponseItems = `[
        {
            "kind": "youtube#video",
            "etag": "dup0Pb5usxFC06Dsb_ET_znf4cg",
            "id": "dQw4w9WgXcQ",
            "snippet": {
                "publishedAt": "2009-10-25T06:57:33Z",
                "channelId": "UCuAXFkgsw1L7xaCfnd5JJOw",
                "title": "Rick Astley - Never Gonna Give You Up (Official Music Video)",
                "description": "The official video for “Never Gonna Give You Up” by Rick Astley\nTaken from the album ‘Whenever You Need Somebody’ – deluxe 2CD and digital deluxe out 6th May 2022 Pre-order here – https://RickAstley.lnk.to/WYNS2022ID\n\n“Never Gonna Give You Up” was a global smash on its release in July 1987, topping the charts in 25 countries including Rick’s native UK and the US Billboard Hot 100.  It also won the Brit Award for Best single in 1988. Stock Aitken and Waterman wrote and produced the track which was the lead-off single and lead track from Rick’s debut LP “Whenever You Need Somebody”.  The album was itself a UK number one and would go on to sell over 15 million copies worldwide.\n\nThe legendary video was directed by Simon West – who later went on to make Hollywood blockbusters such as Con Air, Lara Croft – Tomb Raider and The Expendables 2.  The video passed the 1bn YouTube views milestone on 28 July 2021.\n\nSubscribe to the official Rick Astley YouTube channel: https://RickAstley.lnk.to/YTSubID\n\nFollow Rick Astley:\nFacebook: https://RickAstley.lnk.to/FBFollowID \nTwitter: https://RickAstley.lnk.to/TwitterID \nInstagram: https://RickAstley.lnk.to/InstagramID \nWebsite: https://RickAstley.lnk.to/storeID \nTikTok: https://RickAstley.lnk.to/TikTokID\n\nListen to Rick Astley:\nSpotify: https://RickAstley.lnk.to/SpotifyID \nApple Music: https://RickAstley.lnk.to/AppleMusicID \nAmazon Music: https://RickAstley.lnk.to/AmazonMusicID \nDeezer: https://RickAstley.lnk.to/DeezerID \n\nLyrics:\nWe’re no strangers to love\nYou know the rules and so do I\nA full commitment’s what I’m thinking of\nYou wouldn’t get this from any other guy\n\nI just wanna tell you how I’m feeling\nGotta make you understand\n\nNever gonna give you up\nNever gonna let you down\nNever gonna run around and desert you\nNever gonna make you cry\nNever gonna say goodbye\nNever gonna tell a lie and hurt you\n\nWe’ve known each other for so long\nYour heart’s been aching but you’re too shy to say it\nInside we both know what’s been going on\nWe know the game and we’re gonna play it\n\nAnd if you ask me how I’m feeling\nDon’t tell me you’re too blind to see\n\nNever gonna give you up\nNever gonna let you down\nNever gonna run around and desert you\nNever gonna make you cry\nNever gonna say goodbye\nNever gonna tell a lie and hurt you\n\n#RickAstley #NeverGonnaGiveYouUp #WheneverYouNeedSomebody #OfficialMusicVideo",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/dQw4w9WgXcQ/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/dQw4w9WgXcQ/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/dQw4w9WgXcQ/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    },
                    "standard": {
                        "url": "https://i.ytimg.com/vi/dQw4w9WgXcQ/sddefault.jpg",
                        "width": 640,
                        "height": 480
                    },
                    "maxres": {
                        "url": "https://i.ytimg.com/vi/dQw4w9WgXcQ/maxresdefault.jpg",
                        "width": 1280,
                        "height": 720
                    }
                },
                "channelTitle": "Rick Astley",
                "tags": [
                    "rick astley",
                    "Never Gonna Give You Up",
                    "nggyu",
                    "never gonna give you up lyrics",
                    "rick rolled",
                    "Rick Roll",
                    "rick astley official",
                    "rickrolled",
                    "Fortnite song",
                    "Fortnite event",
                    "Fortnite dance",
                    "fortnite never gonna give you up",
                    "rick roll",
                    "rickrolling",
                    "rick rolling",
                    "never gonna give you up",
                    "80s music",
                    "rick astley new",
                    "animated video",
                    "rickroll",
                    "meme songs",
                    "never gonna give u up lyrics",
                    "Rick Astley 2022",
                    "never gonna let you down",
                    "animated",
                    "rick rolls 2022",
                    "never gonna give you up karaoke"
                ],
                "categoryId": "10",
                "liveBroadcastContent": "none",
                "defaultLanguage": "en",
                "localized": {
                    "title": "Rick Astley - Never Gonna Give You Up (Official Music Video)",
                    "description": "The official video for “Never Gonna Give You Up” by Rick Astley\nTaken from the album ‘Whenever You Need Somebody’ – deluxe 2CD and digital deluxe out 6th May 2022 Pre-order here – https://RickAstley.lnk.to/WYNS2022ID\n\n“Never Gonna Give You Up” was a global smash on its release in July 1987, topping the charts in 25 countries including Rick’s native UK and the US Billboard Hot 100.  It also won the Brit Award for Best single in 1988. Stock Aitken and Waterman wrote and produced the track which was the lead-off single and lead track from Rick’s debut LP “Whenever You Need Somebody”.  The album was itself a UK number one and would go on to sell over 15 million copies worldwide.\n\nThe legendary video was directed by Simon West – who later went on to make Hollywood blockbusters such as Con Air, Lara Croft – Tomb Raider and The Expendables 2.  The video passed the 1bn YouTube views milestone on 28 July 2021.\n\nSubscribe to the official Rick Astley YouTube channel: https://RickAstley.lnk.to/YTSubID\n\nFollow Rick Astley:\nFacebook: https://RickAstley.lnk.to/FBFollowID \nTwitter: https://RickAstley.lnk.to/TwitterID \nInstagram: https://RickAstley.lnk.to/InstagramID \nWebsite: https://RickAstley.lnk.to/storeID \nTikTok: https://RickAstley.lnk.to/TikTokID\n\nListen to Rick Astley:\nSpotify: https://RickAstley.lnk.to/SpotifyID \nApple Music: https://RickAstley.lnk.to/AppleMusicID \nAmazon Music: https://RickAstley.lnk.to/AmazonMusicID \nDeezer: https://RickAstley.lnk.to/DeezerID \n\nLyrics:\nWe’re no strangers to love\nYou know the rules and so do I\nA full commitment’s what I’m thinking of\nYou wouldn’t get this from any other guy\n\nI just wanna tell you how I’m feeling\nGotta make you understand\n\nNever gonna give you up\nNever gonna let you down\nNever gonna run around and desert you\nNever gonna make you cry\nNever gonna say goodbye\nNever gonna tell a lie and hurt you\n\nWe’ve known each other for so long\nYour heart’s been aching but you’re too shy to say it\nInside we both know what’s been going on\nWe know the game and we’re gonna play it\n\nAnd if you ask me how I’m feeling\nDon’t tell me you’re too blind to see\n\nNever gonna give you up\nNever gonna let you down\nNever gonna run around and desert you\nNever gonna make you cry\nNever gonna say goodbye\nNever gonna tell a lie and hurt you\n\n#RickAstley #NeverGonnaGiveYouUp #WheneverYouNeedSomebody #OfficialMusicVideo"
                },
                "defaultAudioLanguage": "en"
            },
            "contentDetails": {
                "duration": "PT3M33S",
                "dimension": "2d",
                "definition": "hd",
                "caption": "false",
                "licensedContent": true,
                "contentRating": {},
                "projection": "rectangular"
            },
            "statistics": {
                "viewCount": "1299408727",
                "likeCount": "15432598",
                "favoriteCount": "0",
                "commentCount": "2217355"
            }
        }
    ]`

var VideoSearchResponse = fmt.Sprintf(`{
    "kind": "youtube#searchListResponse",
    "etag": "5sxu8HTLbhWJvSXX8EBQI2z6gz0",
    "nextPageToken": "CDIQAA",
	"prevPageToken": "",
    "regionCode": "DE",
    "pageInfo": {
        "totalResults": 1000000,
        "resultsPerPage": 50
    },
    "items": %s
}`, VideoSearchResponseItems)

var VideoSearchResponseItems = `[
        {
            "kind": "youtube#searchResult",
            "etag": "4AI2Sz2ZqtoKsmjsuUsI8HsJ3as",
            "id": {
                "kind": "youtube#video",
                "videoId": "dQw4w9WgXcQ"
            },
            "snippet": {
                "publishedAt": "2009-10-25T06:57:33Z",
                "channelId": "UCuAXFkgsw1L7xaCfnd5JJOw",
                "title": "Rick Astley - Never Gonna Give You Up (Official Music Video)",
                "description": "The official video for “Never Gonna Give You Up” by Rick Astley Taken from the album 'Whenever You Need Somebody' – deluxe ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/dQw4w9WgXcQ/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/dQw4w9WgXcQ/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/dQw4w9WgXcQ/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "Rick Astley",
                "liveBroadcastContent": "none",
                "publishTime": "2009-10-25T06:57:33Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "jVWWJX9zzbe_Po3BvYiNRiEOTO8",
            "id": {
                "kind": "youtube#video",
                "videoId": "34Ig3X59_qA"
            },
            "snippet": {
                "publishedAt": "2020-10-07T18:23:37Z",
                "channelId": "UCvR2R7j218tzejtTsb_X6Rw",
                "title": "Never Gonna Give You Up - Rick Astley (Lyrics) 🎵",
                "description": "Rick Astley - Never Gonna Give You Up (Lyrics) Submit your music here: https://www.unitedframes.tv/dopelyrics For more ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/34Ig3X59_qA/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/34Ig3X59_qA/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/34Ig3X59_qA/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "DopeLyrics",
                "liveBroadcastContent": "none",
                "publishTime": "2020-10-07T18:23:37Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "NmsdZjVOMFnGyeSQuAJUhNk8cF0",
            "id": {
                "kind": "youtube#video",
                "videoId": "GtL1huin9EE"
            },
            "snippet": {
                "publishedAt": "2022-08-15T16:37:13Z",
                "channelId": "UCJjjxrjyjMMKFzq23xsaK-Q",
                "title": "InsurAAAnce &amp; Rick Astley Never Gonna Give You Up",
                "description": "Our legendary service is never gonna let you down. Visit https://csaa-insurance.aaa.com/rickroll and you can start saving today on ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/GtL1huin9EE/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/GtL1huin9EE/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/GtL1huin9EE/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "CSAA Insurance Group, a AAA Insurer",
                "liveBroadcastContent": "none",
                "publishTime": "2022-08-15T16:37:13Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "_m5bq4RTucKfcMQVVeDVXRBU0QA",
            "id": {
                "kind": "youtube#video",
                "videoId": "o-YBDTqX_ZU"
            },
            "snippet": {
                "publishedAt": "2021-02-19T13:42:40Z",
                "channelId": "UC2eEHetmyPERpOQzJzEdZpA",
                "title": "Rick Astley - Never Gonna Give You Up (Remastered 4K 60fps,AI)",
                "description": "(Digital Remaster) The original file was restored and processed at 480p quality in 2160р60fps. DONATION ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/o-YBDTqX_ZU/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/o-YBDTqX_ZU/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/o-YBDTqX_ZU/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "MusRest",
                "liveBroadcastContent": "none",
                "publishTime": "2021-02-19T13:42:40Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "nSTM4rci-MagKYDO36NYyY4IFKE",
            "id": {
                "kind": "youtube#video",
                "videoId": "VbUuB1aN2DA"
            },
            "snippet": {
                "publishedAt": "2019-06-15T20:29:39Z",
                "channelId": "UCnULp16WtSX404Du_Nlmcxw",
                "title": "Rick Astley - Never Gonna Give You Up - Live at The Isle of Wight Festival 2019",
                "description": "",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/VbUuB1aN2DA/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/VbUuB1aN2DA/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/VbUuB1aN2DA/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "isleofwightfestival",
                "liveBroadcastContent": "none",
                "publishTime": "2019-06-15T20:29:39Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "gnL6M3E0DPkeIe7sT0FSd1FmNq8",
            "id": {
                "kind": "youtube#video",
                "videoId": "xm3YgoEiEDc"
            },
            "snippet": {
                "publishedAt": "2021-07-22T20:24:54Z",
                "channelId": "UCZil07Hq4ESfFOmbbe-2Umg",
                "title": "10 Hours Rick Astley Never Gonna Give You Up",
                "description": "10 Hours Rick Astley Never Gonna Give You Up.",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/xm3YgoEiEDc/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/xm3YgoEiEDc/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/xm3YgoEiEDc/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "10 Hours",
                "liveBroadcastContent": "none",
                "publishTime": "2021-07-22T20:24:54Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "8jpUYT97GdWfoFY0Da3LZm9aZdU",
            "id": {
                "kind": "youtube#video",
                "videoId": "3BFTio5296w"
            },
            "snippet": {
                "publishedAt": "2022-05-19T10:03:26Z",
                "channelId": "UCwZEU0wAwIyZb4x5G_KJp2w",
                "title": "Never Gonna Give You Up (2022 - Remaster)",
                "description": "Provided to YouTube by BMG Rights Management (UK) Ltd. Never Gonna Give You Up (2022 - Remaster) · Rick Astley Whenever ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/3BFTio5296w/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/3BFTio5296w/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/3BFTio5296w/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "Rick Astley - Topic",
                "liveBroadcastContent": "none",
                "publishTime": "2022-05-19T10:03:26Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "qQn0UNF03pB9j-QOnhRpMf8YqVA",
            "id": {
                "kind": "youtube#video",
                "videoId": "dJRsWJqDjFE"
            },
            "snippet": {
                "publishedAt": "2018-04-18T18:03:24Z",
                "channelId": "UC0Kte42xsXIT80UzG1jotiw",
                "title": "Choir! Choir! Choir! / Rick Astley - Never Gonna Give You Up!!!",
                "description": "In March 2018, we made a video asking pop legend Rick Astley to come sing with us. HE SHOWED UP two weeks later!! It was ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/dJRsWJqDjFE/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/dJRsWJqDjFE/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/dJRsWJqDjFE/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "Choir! Choir! Choir!",
                "liveBroadcastContent": "none",
                "publishTime": "2018-04-18T18:03:24Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "qzp4jOR9Y770JH7PoMR6xrct4jo",
            "id": {
                "kind": "youtube#video",
                "videoId": "oADU2PIzhD0"
            },
            "snippet": {
                "publishedAt": "2022-10-08T15:00:38Z",
                "channelId": "UCn8zNIfYAQNdrFRrr8oibKw",
                "title": "The Legendary Song That Became The Rick Roll | The Story Of",
                "description": "Nearly everyone has been RickRoll'd. But few know the real story of the artist behind the 1987 hit “Never Gonna Give You Up” nor ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/oADU2PIzhD0/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/oADU2PIzhD0/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/oADU2PIzhD0/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "VICE",
                "liveBroadcastContent": "none",
                "publishTime": "2022-10-08T15:00:38Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "pgx8IYQ2j6iMlucGXYJZmk1d3-g",
            "id": {
                "kind": "youtube#video",
                "videoId": "WmttqfVT830"
            },
            "snippet": {
                "publishedAt": "2018-10-24T14:00:07Z",
                "channelId": "UC0_LJ1oJtR5keARJWupA3_g",
                "title": "Rick Astley - Never Gonna Give You Up (Jeanie Schultheiß) | The Voice of Germany | Blind Audition",
                "description": "Mit so viel Gefühl haben wir den 80er-Hit von Rick Astley \"Never Gonna Give You Up\" noch nie gehört! Jeanie Schultheiß ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/WmttqfVT830/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/WmttqfVT830/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/WmttqfVT830/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "The Voice of Germany - Offiziell",
                "liveBroadcastContent": "none",
                "publishTime": "2018-10-24T14:00:07Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "sSU5h1BAl_RH_zBwL80d0iaX6WI",
            "id": {
                "kind": "youtube#video",
                "videoId": "NfgZZfI_j5M"
            },
            "snippet": {
                "publishedAt": "2021-07-30T17:35:27Z",
                "channelId": "UC4YCtfBDKm5cunRfkYbASEg",
                "title": "Rick Astley&#39;s &quot;Never Gonna Give You Up&quot; in different languages",
                "description": "\"Never Gonna Give You Up\" in 12 different languages #RickAstley #InDifferentLanguages #NeverGonnaGiveYouUp rickrolled,the ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/NfgZZfI_j5M/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/NfgZZfI_j5M/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/NfgZZfI_j5M/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "In Different Languages",
                "liveBroadcastContent": "none",
                "publishTime": "2021-07-30T17:35:27Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "NQ76Fft_vRVte6SnGnL55D32iNU",
            "id": {
                "kind": "youtube#video",
                "videoId": "-PB_XHb93bk"
            },
            "snippet": {
                "publishedAt": "2022-02-17T19:19:28Z",
                "channelId": "UCWzTTaxhiI0KEqZS4FRVMzw",
                "title": "Memories x Bad And Boujee x Never Gonna Give You Up (Tiktok Mashup) Tibodd Migos x David Guetta",
                "description": "Hi! I perfectly extend popular TikTok songs. Follow our spotify playlist: https://sptfy.com/latenightphonk.",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/-PB_XHb93bk/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/-PB_XHb93bk/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/-PB_XHb93bk/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "TikTok Extended",
                "liveBroadcastContent": "none",
                "publishTime": "2022-02-17T19:19:28Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "PAKrJgC4a1MhjSCgJzUW0qzA3uQ",
            "id": {
                "kind": "youtube#video",
                "videoId": "6_b7RDuLwcI"
            },
            "snippet": {
                "publishedAt": "2008-07-02T18:28:59Z",
                "channelId": "UCuoFoUKfg2XsApaZAvdSvPA",
                "title": "Rick Astley Never gonna give you up lyrics!!!",
                "description": "the rick rolled thing was an inside joke thing. not many people know what it means. but yet alot do!",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/6_b7RDuLwcI/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/6_b7RDuLwcI/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/6_b7RDuLwcI/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "Jaysean",
                "liveBroadcastContent": "none",
                "publishTime": "2008-07-02T18:28:59Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "2T82p8Rg8EEbuniXGI1B-XuqmVU",
            "id": {
                "kind": "youtube#video",
                "videoId": "IdkCEioCp24"
            },
            "snippet": {
                "publishedAt": "2017-09-20T00:38:07Z",
                "channelId": "UCy2BAnjgfPsi6abufy7ysmg",
                "title": "Foo Fighters With Rick Astley - Never Gonna Give You Up  - London O2 Arena 19 September 2017",
                "description": "Foo Fighters With Rick Astley - Never Gonna Give You Up - London O2 Arena 19 September 2017.",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/IdkCEioCp24/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/IdkCEioCp24/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/IdkCEioCp24/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "GotsomePearlJam",
                "liveBroadcastContent": "none",
                "publishTime": "2017-09-20T00:38:07Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "-YmzNecasb4OwmDfzln6eEloSfk",
            "id": {
                "kind": "youtube#video",
                "videoId": "rTgj1HxmUbg"
            },
            "snippet": {
                "publishedAt": "2020-10-02T03:00:12Z",
                "channelId": "UC8PzIXMp5gkDh82PfKbbHmA",
                "title": "Rick Astley Never gonna give you up 1 hour seamless loop",
                "description": "the second most ultimate rickroll 12-hour version:https://youtu.be/zL19uMsnpSU.",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/rTgj1HxmUbg/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/rTgj1HxmUbg/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/rTgj1HxmUbg/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "cameron barnett",
                "liveBroadcastContent": "none",
                "publishTime": "2020-10-02T03:00:12Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "KHq5gUGMhZzyxBk9RVe6J71cEFA",
            "id": {
                "kind": "youtube#video",
                "videoId": "DsC8jQXRbQE"
            },
            "snippet": {
                "publishedAt": "2010-12-30T07:01:49Z",
                "channelId": "UCgTtefAKukDNOs0G4tT44aQ",
                "title": "Family Guy - Never Gonna Give You Up",
                "description": "Brian sing a song from the future . S05E18.",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/DsC8jQXRbQE/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/DsC8jQXRbQE/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/DsC8jQXRbQE/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "Rep Jay",
                "liveBroadcastContent": "none",
                "publishTime": "2010-12-30T07:01:49Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "mWegvchKyi1V2gMwzz3JT-9PmUc",
            "id": {
                "kind": "youtube#video",
                "videoId": "ur560pZKRfg"
            },
            "snippet": {
                "publishedAt": "2020-10-27T16:30:29Z",
                "channelId": "UCP2-S6-M9ZvlY8t7cRn4O6A",
                "title": "Google Translate Sings: &quot;Never Gonna Give You Up&quot; by Rick Astley",
                "description": "SUBSCRIBE: http://bit.ly/sub2Malinda SUBSCRIBE TO TOM: https://www.youtube.com/channel/UCmRFGJTyduULs1WhcNwYIuQ ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/ur560pZKRfg/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/ur560pZKRfg/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/ur560pZKRfg/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "Twisted Translations",
                "liveBroadcastContent": "none",
                "publishTime": "2020-10-27T16:30:29Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "RuKUzqk2ksGRKv6AM6HdDj9bDn8",
            "id": {
                "kind": "youtube#video",
                "videoId": "la-GFyRzIRA"
            },
            "snippet": {
                "publishedAt": "2022-04-14T15:38:41Z",
                "channelId": "UCstYrD4DvY2KNWiO1E4Ay_A",
                "title": "Never Gonna Give You Up - 1 Hour Version - Rick Astley (Lyrics)",
                "description": "Don't forget to subscribe if you enjoyed the video!! It means the world to me to see my subscriber count go up, even by one.",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/la-GFyRzIRA/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/la-GFyRzIRA/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/la-GFyRzIRA/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "Daily Dose Of Songs",
                "liveBroadcastContent": "none",
                "publishTime": "2022-04-14T15:38:41Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "MR2v8AuRPROOhwW-NOHcSDpYT9k",
            "id": {
                "kind": "youtube#video",
                "videoId": "8oE5Z2GLhNc"
            },
            "snippet": {
                "publishedAt": "2022-06-10T17:30:11Z",
                "channelId": "UCh34OFHpOjQSRZWhvBpY39w",
                "title": "Yung Gravy - Betty (Get Money) (Official Music Video)",
                "description": "Betty” (Get Money) out now on Spotify, Apple Music, and all your favorite streaming services: https://YungGravy.lnk.to/betty ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/8oE5Z2GLhNc/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/8oE5Z2GLhNc/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/8oE5Z2GLhNc/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "YungGravyVEVO",
                "liveBroadcastContent": "none",
                "publishTime": "2022-06-10T17:30:11Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "Tlxb774dJ_eH690_D5nN9n0JQPA",
            "id": {
                "kind": "youtube#video",
                "videoId": "v_nRCCD0j-k"
            },
            "snippet": {
                "publishedAt": "2021-04-02T17:00:14Z",
                "channelId": "UCTFuNYrqAcsmSjgqYMvxOqw",
                "title": "Home Free - Never Gonna Give You Up (Official Music Video)",
                "description": "YOU'VE BEEN RICKROLLED ▻ Thanks for watching y'all! If you liked this, then you'll love seeing us live ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/v_nRCCD0j-k/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/v_nRCCD0j-k/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/v_nRCCD0j-k/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "Home Free",
                "liveBroadcastContent": "none",
                "publishTime": "2021-04-02T17:00:14Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "Q_1hyNpZuFyKw_ayHK2FY4PatOA",
            "id": {
                "kind": "youtube#video",
                "videoId": "zL19uMsnpSU"
            },
            "snippet": {
                "publishedAt": "2021-06-13T09:04:43Z",
                "channelId": "UC8PzIXMp5gkDh82PfKbbHmA",
                "title": "Rick Astley Never gonna give you up 12 hour seamless loop",
                "description": "the ultimate rickroll. remade my most popular video to celebrate 1000 subscribers 12 times as long and now with video, enjoy.",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/zL19uMsnpSU/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/zL19uMsnpSU/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/zL19uMsnpSU/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "cameron barnett",
                "liveBroadcastContent": "none",
                "publishTime": "2021-06-13T09:04:43Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "RP1QHBeqL8X8g44MlLjaJPE3bkE",
            "id": {
                "kind": "youtube#video",
                "videoId": "a9WHZ5M8I8w"
            },
            "snippet": {
                "publishedAt": "2018-02-07T18:00:00Z",
                "channelId": "UCTqDzPyS7dDBxjywNROdAUg",
                "title": "Rick Astley - Never Gonna Give You Up - Festival de Viña del Mar 2016 HD",
                "description": "Rick Astley - Never Gonna Give You Up - Festival de Viña del Mar 2016 HD No olvides suscribirte y activar la campanita para ser ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/a9WHZ5M8I8w/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/a9WHZ5M8I8w/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/a9WHZ5M8I8w/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "FESTIVALDEVINACHILE",
                "liveBroadcastContent": "none",
                "publishTime": "2018-02-07T18:00:00Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "gh3ZwHycc3oPcjHbizQiZKwbYwE",
            "id": {
                "kind": "youtube#video",
                "videoId": "06kbuUQlvbw"
            },
            "snippet": {
                "publishedAt": "2021-10-04T14:53:55Z",
                "channelId": "UCzlSGCgkLcpV-aajtLcjhEw",
                "title": "YouTubers getting Rick rolled compilation! ( PewDiePie, MrBeast, Dream and more!)",
                "description": "Music in the beginning of the video Music: BigCarTheft Musician: Jason Shaw URL: https://audionautix.com License: ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/06kbuUQlvbw/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/06kbuUQlvbw/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/06kbuUQlvbw/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "XIX Fr",
                "liveBroadcastContent": "none",
                "publishTime": "2021-10-04T14:53:55Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "XqKiwJAbK1dWedhniiTbetgSFdo",
            "id": {
                "kind": "youtube#video",
                "videoId": "5lUYJpoNt4g"
            },
            "snippet": {
                "publishedAt": "2021-07-02T07:33:32Z",
                "channelId": "UCebNYMaQBePWH282FiYQBXQ",
                "title": "Never gonna give you up lyrics",
                "description": "",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/5lUYJpoNt4g/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/5lUYJpoNt4g/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/5lUYJpoNt4g/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "Lyrics studio",
                "liveBroadcastContent": "none",
                "publishTime": "2021-07-02T07:33:32Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "qzvVJ_F1xGg8j0EaPLK0F2ampHc",
            "id": {
                "kind": "youtube#video",
                "videoId": "jcZ2A8G_PrA"
            },
            "snippet": {
                "publishedAt": "2022-08-18T05:36:37Z",
                "channelId": "UCELAPsjNdBBFOMfd3nyavLA",
                "title": "&quot;Never Gonna Give You Up&quot; (1987 VS 2022) Side-by-side Comparison | Rick Astley &amp; InsurAAAnce",
                "description": "rickroll #rickastley #nevergonnagiveyouup #InsurAAAnce \"Never Gonna Give You Up\" - Rick Astley (1987) VS InsurAAAnce ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/jcZ2A8G_PrA/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/jcZ2A8G_PrA/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/jcZ2A8G_PrA/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "阿湧",
                "liveBroadcastContent": "none",
                "publishTime": "2022-08-18T05:36:37Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "WrOH-GEf1xEl3bOdQdnJKgrAujo",
            "id": {
                "kind": "youtube#video",
                "videoId": "FzWQg7hH80g"
            },
            "snippet": {
                "publishedAt": "2022-08-27T15:06:57Z",
                "channelId": "UCqwIRJDpvqUFGK07PoHuLvA",
                "title": "Rick Astley - Never Gonna Give You Up nueva version 2022",
                "description": "Rick Astley - Never Gonna Give You Up nueva version 2022 super espectacular.",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/FzWQg7hH80g/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/FzWQg7hH80g/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/FzWQg7hH80g/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "MINITECAS TV",
                "liveBroadcastContent": "none",
                "publishTime": "2022-08-27T15:06:57Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "9ytY1Iw4tJHoBKTG88SRZB7kdVs",
            "id": {
                "kind": "youtube#video",
                "videoId": "-DdkOB0htO0"
            },
            "snippet": {
                "publishedAt": "2022-05-20T05:58:44Z",
                "channelId": "UCwEwTwfIyFOLONqzwZWzpXQ",
                "title": "Rick Astley Live 2022 🡆 Never Gonna Give You Up 🡄 May 20 ⬘ Houston, TX",
                "description": "watch the show: https://www.youtube.com/playlist?list=PL0oz3aZg2QY6CjI8oy2PbGLMe9F0LIhtp Rick waves at me - 3:04 shot on ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/-DdkOB0htO0/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/-DdkOB0htO0/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/-DdkOB0htO0/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "DeadMike.com",
                "liveBroadcastContent": "none",
                "publishTime": "2022-05-20T05:58:44Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "0hy0xW4xvIZ5wyj14sjvjKwl-Rk",
            "id": {
                "kind": "youtube#video",
                "videoId": "OiZlXOAOLLw"
            },
            "snippet": {
                "publishedAt": "2018-09-10T13:15:14Z",
                "channelId": "UCZtDUmC3W7j25XHZWFT_XgQ",
                "title": "Kylie Minogue &amp; Rick Astley - I Should Be So Lucky /Never Gonna Give You Up (Hyde Park 2018)",
                "description": "Rick Astley joins Kylie on stage at Radio 2's Festival in a Day to perform Never Gonna Give You Up and I Should Be So Lucky.",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/OiZlXOAOLLw/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/OiZlXOAOLLw/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/OiZlXOAOLLw/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "BBC Music",
                "liveBroadcastContent": "none",
                "publishTime": "2018-09-10T13:15:14Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "vkAkIQq5V5gUvYH25yaVUa48Ibo",
            "id": {
                "kind": "youtube#video",
                "videoId": "tVwUn4r4Byo"
            },
            "snippet": {
                "publishedAt": "2021-04-24T12:56:15Z",
                "channelId": "UCKIUy1Qg9rgQFWJoYQtYZ0A",
                "title": "Never gonna give you up - wreck it ralph version",
                "description": "",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/tVwUn4r4Byo/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/tVwUn4r4Byo/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/tVwUn4r4Byo/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "SlyBur",
                "liveBroadcastContent": "none",
                "publishTime": "2021-04-24T12:56:15Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "IituGMxsskmZJZMuYZFpWMY4jgk",
            "id": {
                "kind": "youtube#video",
                "videoId": "vxUxkCQR_EM"
            },
            "snippet": {
                "publishedAt": "2018-10-26T16:25:08Z",
                "channelId": "UCCj956IF62FbT7Gouszaj9w",
                "title": "Rick’s WORST Never Gonna Give You Up reactions - BBC",
                "description": "Subscribe and to the BBC https://bit.ly/BBCYouTubeSub Watch the BBC first on iPlayer https://bbc.in/iPlayer-Home How ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/vxUxkCQR_EM/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/vxUxkCQR_EM/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/vxUxkCQR_EM/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "BBC",
                "liveBroadcastContent": "none",
                "publishTime": "2018-10-26T16:25:08Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "a6JcfRGlEJT8ivUbSoxen9vxa7s",
            "id": {
                "kind": "youtube#video",
                "videoId": "5rDjqsiDT9Y"
            },
            "snippet": {
                "publishedAt": "2022-08-19T00:00:49Z",
                "channelId": "UCoBCss5PoyJ8A9HSJRP9IzQ",
                "title": "RICK ASTLEY recreó su exitazo 35 años despues",
                "description": "RICK ASTLEY recreó su exitazo 35 años despues y se volvió tendencia.",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/5rDjqsiDT9Y/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/5rDjqsiDT9Y/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/5rDjqsiDT9Y/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "La Maquina del Tiempo",
                "liveBroadcastContent": "none",
                "publishTime": "2022-08-19T00:00:49Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "AfQkPNfhtAk7wjENl4ljjCjStuk",
            "id": {
                "kind": "youtube#video",
                "videoId": "PayvWj2piKg"
            },
            "snippet": {
                "publishedAt": "2020-08-12T15:30:43Z",
                "channelId": "UCYQT13AtrJC0gsM1far_zJg",
                "title": "Rick Astley - Give up",
                "description": "BUY FLYINGKITTY PLUSHIE HERE: https://makeship.com/products/flyingkitty BUY FLYINGKITTY PLUSHIE HERE: ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/PayvWj2piKg/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/PayvWj2piKg/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/PayvWj2piKg/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "FlyingKitty",
                "liveBroadcastContent": "none",
                "publishTime": "2020-08-12T15:30:43Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "iSGcFxINSpFUZL_fieMLCtAYLjg",
            "id": {
                "kind": "youtube#video",
                "videoId": "qPJuzB8CL8o"
            },
            "snippet": {
                "publishedAt": "2022-03-31T16:46:48Z",
                "channelId": "UCy3jC5OdSSrrkbe_496i6gQ",
                "title": "If Blink 182 Wrote &#39;Never Gonna Give You Up&#39;",
                "description": "Thanks to http://www.zZounds.com for sponsoring the channel! Go check 'em out. The ultimate internet inside joke is now in pop ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/qPJuzB8CL8o/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/qPJuzB8CL8o/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/qPJuzB8CL8o/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "Alex Melton",
                "liveBroadcastContent": "none",
                "publishTime": "2022-03-31T16:46:48Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "9iuFoFJJzqKpW3BL3y-sGMp3Fe0",
            "id": {
                "kind": "youtube#video",
                "videoId": "iJgNpm8cTE8"
            },
            "snippet": {
                "publishedAt": "2020-05-01T14:03:13Z",
                "channelId": "UCxNNwgtDWEoDmq_rqZeXISQ",
                "title": "Never Gonna Give You Up, but an AI creates more of the song",
                "description": "So this neural network is called Jukebox, by OpenAI, which basically generates songs that don't exist. Check out the links below if ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/iJgNpm8cTE8/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/iJgNpm8cTE8/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/iJgNpm8cTE8/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "Lil'Alien",
                "liveBroadcastContent": "none",
                "publishTime": "2020-05-01T14:03:13Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "YQrE_8RqOGAS_Q7FSqy82MzaV6g",
            "id": {
                "kind": "youtube#video",
                "videoId": "ci6ZtPAN0PM"
            },
            "snippet": {
                "publishedAt": "2022-06-18T14:06:06Z",
                "channelId": "UCmSp4bDxS9R0jpeZEvkut2g",
                "title": "RickRolled by an Ad...",
                "description": "This is big brain time.",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/ci6ZtPAN0PM/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/ci6ZtPAN0PM/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/ci6ZtPAN0PM/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "Beluga",
                "liveBroadcastContent": "none",
                "publishTime": "2022-06-18T14:06:06Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "04qbA5sOam4dojGDAuNokQrDu30",
            "id": {
                "kind": "youtube#video",
                "videoId": "sG2oz-hfnZU"
            },
            "snippet": {
                "publishedAt": "2022-08-06T18:00:22Z",
                "channelId": "UCFV5crW1yQWGw1QmjVCzmBQ",
                "title": "Rick Astley going back in time… (THE MOVIE)",
                "description": "I'm not adding credits it's too much but if you do want the credits go to every single one of my Rick Astley videos￼ ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/sG2oz-hfnZU/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/sG2oz-hfnZU/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/sG2oz-hfnZU/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "Flop Meme",
                "liveBroadcastContent": "none",
                "publishTime": "2022-08-06T18:00:22Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "HzL2LWTMAqXaIX3JiN2euk38TPs",
            "id": {
                "kind": "youtube#video",
                "videoId": "KyElxl_j4Wc"
            },
            "snippet": {
                "publishedAt": "2021-11-26T13:00:11Z",
                "channelId": "UCdXHc-g-qmrGpDFuSEwWbAg",
                "title": "Rickroll Minecraft Note Block (Rick Astley - Never Gonna Give You Up)",
                "description": "Thank you for watching my videos. I hope you enjoy. --------------------------------------------------------------- Tools : Kinemaster (Video ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/KyElxl_j4Wc/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/KyElxl_j4Wc/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/KyElxl_j4Wc/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "Mr Aldi",
                "liveBroadcastContent": "none",
                "publishTime": "2021-11-26T13:00:11Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "ruK3cDUZguMM_sDFfVG8teflKbA",
            "id": {
                "kind": "youtube#video",
                "videoId": "pzYpHXCumII"
            },
            "snippet": {
                "publishedAt": "2009-11-15T10:23:07Z",
                "channelId": "UCeA38jYPEzkE4TbUHiC5CLw",
                "title": "Lisa Stansfield - Never, Never Gonna Give You Up (Video)",
                "description": "Music video by Lisa Stansfield performing Never, Never Gonna Give You Up. (c) 1997 Sony Music Entertainment UK Limited.",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/pzYpHXCumII/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/pzYpHXCumII/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/pzYpHXCumII/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "LisaStansfieldVEVO",
                "liveBroadcastContent": "none",
                "publishTime": "2009-11-15T10:23:07Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "jwh4P5uNdw4Nwfc6A0KYF8PocrE",
            "id": {
                "kind": "youtube#video",
                "videoId": "TMBN8hj5Jy0"
            },
            "snippet": {
                "publishedAt": "2018-05-05T19:44:54Z",
                "channelId": "UC_RXAvEdK0a8HjP4jkTCCaA",
                "title": "Rick Astley - Never Gonna Give You Up (Live2016) Tradução",
                "description": "Um show , simplesmente Maravilhoso ! Uma das melhores vozes de todos os tempos . Rick Astley - Never Gonna Give You Up ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/TMBN8hj5Jy0/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/TMBN8hj5Jy0/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/TMBN8hj5Jy0/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "KALUNGA - I Love Musiic",
                "liveBroadcastContent": "none",
                "publishTime": "2018-05-05T19:44:54Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "2V2bNlBxNUsEbzU3oyz90iFbUXQ",
            "id": {
                "kind": "youtube#video",
                "videoId": "cErgMJSgpv0"
            },
            "snippet": {
                "publishedAt": "2021-07-03T17:58:44Z",
                "channelId": "UCbVcb9puAsOhXBT2_XPFf-A",
                "title": "Rīċa Ēastlēah - Never gonna give you up Cover In Old English (BARDCORE)",
                "description": "You should have been here for the premiere and got Rīċhlāfinċelettan. XD (Big thanks to Hroðgar Gliƿþeof for OE-fying the ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/cErgMJSgpv0/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/cErgMJSgpv0/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/cErgMJSgpv0/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "the_miracle_aligner",
                "liveBroadcastContent": "none",
                "publishTime": "2021-07-03T17:58:44Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "klVWUNQEmrEXOq6GeRWWLbBsRBg",
            "id": {
                "kind": "youtube#video",
                "videoId": "eE4UvgqRE2c"
            },
            "snippet": {
                "publishedAt": "2022-08-31T23:39:10Z",
                "channelId": "UCqIQu-eKuPIqIrhT6kWLKGg",
                "title": "Never Gonna Give You Up - Missasinfonia Cover (Extendido)",
                "description": "Letra No es nuevo este amor ambos sabemos bien que hacer un compromiso serio quiero yo y eso cualquiera no te lo va a ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/eE4UvgqRE2c/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/eE4UvgqRE2c/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/eE4UvgqRE2c/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "M4tt_Ch",
                "liveBroadcastContent": "none",
                "publishTime": "2022-08-31T23:39:10Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "7TeWQ_Bql82kNFpRnE8zN-sXJ2M",
            "id": {
                "kind": "youtube#video",
                "videoId": "NCknXQd1f1w"
            },
            "snippet": {
                "publishedAt": "2022-08-13T12:37:25Z",
                "channelId": "UCGlrmkGiIu-FFAxcifnng4g",
                "title": "I Rickrolled TommyInnit&#39;s Minecraft Talent Show",
                "description": "I got invited to TommyInnit's talent show! It was HILARIOUS! One might even call it Minecraft's Funniest YouTuber talent show.",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/NCknXQd1f1w/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/NCknXQd1f1w/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/NCknXQd1f1w/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "iTMG",
                "liveBroadcastContent": "none",
                "publishTime": "2022-08-13T12:37:25Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "x8tPiNcnwoZhHgNstO_4-bN1qy0",
            "id": {
                "kind": "youtube#video",
                "videoId": "DLzxrzFCyOs"
            },
            "snippet": {
                "publishedAt": "2015-04-12T00:23:38Z",
                "channelId": "UCLNd5EtH77IyN1frExzwPRQ",
                "title": "Rick Astley - Never Gonna Give You Up [HQ]",
                "description": "Artist: Rick Astley Title: Never Gonna Give You Up Difference with original: nothing but the quality of audio and video. all rights ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/DLzxrzFCyOs/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/DLzxrzFCyOs/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/DLzxrzFCyOs/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "AllKindsOfStuff",
                "liveBroadcastContent": "none",
                "publishTime": "2015-04-12T00:23:38Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "eNqpoSJG3T0kdYq9ahFc5djVNO4",
            "id": {
                "kind": "youtube#video",
                "videoId": "nc5oOlaGPlA"
            },
            "snippet": {
                "publishedAt": "2022-06-24T06:15:01Z",
                "channelId": "UCsklcAeg3tlI6rpYsjHCSwQ",
                "title": "Never Gonna Make You Ghen",
                "description": "biết là sẽ bị bản quyền nhạc nhưng mà không còn idea nào khác nên vẫn phải đăng:( Bản Remix by Nòng Nọc Music: ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/nc5oOlaGPlA/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/nc5oOlaGPlA/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/nc5oOlaGPlA/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "Loly1478",
                "liveBroadcastContent": "none",
                "publishTime": "2022-06-24T06:15:01Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "fEX3b-MsCC1T6O37BRuqa6_OUYs",
            "id": {
                "kind": "youtube#video",
                "videoId": "igWazH_Babw"
            },
            "snippet": {
                "publishedAt": "2009-03-02T00:02:07Z",
                "channelId": "UCBEsiNgQv4nw-5_rToxOj6w",
                "title": "Ashley Tisdale - Never Gonna Give You Up - Degree Girl Cover -",
                "description": "Title: Never Gonna Give You Up Artist: Ashley Tisdale (Main Artist is Rick Astley) for her \"Degree Girl\" Campaign My opinion is that ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/igWazH_Babw/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/igWazH_Babw/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/igWazH_Babw/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "Chantel",
                "liveBroadcastContent": "none",
                "publishTime": "2009-03-02T00:02:07Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "cKIklb2IlmdzetfPlamfafZKZRs",
            "id": {
                "kind": "youtube#video",
                "videoId": "nBrF-cJNZc8"
            },
            "snippet": {
                "publishedAt": "2021-08-19T18:33:12Z",
                "channelId": "UCAi_uNeDWRXj8C8yw358gWw",
                "title": "If NEVER GONNA GIVE YOU UP Was The Hardest Song In The World",
                "description": "Visit https://betterhelp.com/charles to get 10% off your first month, thanks so much to BetterHelp for sponsoring this video! Let me ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/nBrF-cJNZc8/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/nBrF-cJNZc8/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/nBrF-cJNZc8/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "CharlesBerthoud",
                "liveBroadcastContent": "none",
                "publishTime": "2021-08-19T18:33:12Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "wVED-k5kKC5yaFJxnCss-WYVqGQ",
            "id": {
                "kind": "youtube#video",
                "videoId": "oT3mCybbhf0"
            },
            "snippet": {
                "publishedAt": "2014-02-18T21:21:43Z",
                "channelId": "UC6hhffEdx4fruP8-0MzqHPQ",
                "title": "AVICII &amp; RICK ASTLEY - Never Gonna Wake You Up (NilsOfficial Mashup)",
                "description": "Check out the first original music video I made with my friends : https://www.youtube.com/watch?v=Axf0n5HjLYA Audio Mashup: ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/oT3mCybbhf0/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/oT3mCybbhf0/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/oT3mCybbhf0/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "NilsOfficial",
                "liveBroadcastContent": "none",
                "publishTime": "2014-02-18T21:21:43Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "Pyw0Nj09cK6pJrvlk-yKbpsHklg",
            "id": {
                "kind": "youtube#video",
                "videoId": "eM9FjdBvhOU"
            },
            "snippet": {
                "publishedAt": "2022-08-20T23:35:54Z",
                "channelId": "UC2Zn6cEkFvFnj3X71H8zMgg",
                "title": "IF YOU ARE SUBBING TO THIS CHANNEL SUB TO TheRealStrims (My Main) :(",
                "description": "",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/eM9FjdBvhOU/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/eM9FjdBvhOU/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/eM9FjdBvhOU/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "TheRealStrims Alt",
                "liveBroadcastContent": "none",
                "publishTime": "2022-08-20T23:35:54Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "8PCKFkWH4IqS4aa8Z9QxYOx92iE",
            "id": {
                "kind": "youtube#video",
                "videoId": "yPYZpwSpKmA"
            },
            "snippet": {
                "publishedAt": "2009-10-25T18:56:16Z",
                "channelId": "UCuAXFkgsw1L7xaCfnd5JJOw",
                "title": "Rick Astley - Together Forever (Official Video) [Remastered in 4K]",
                "description": "The official video for “Together Forever” by Rick Astley Taken from the album 'Whenever You Need Somebody' (2022 Remaster) ...",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/yPYZpwSpKmA/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/yPYZpwSpKmA/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/yPYZpwSpKmA/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "Rick Astley",
                "liveBroadcastContent": "none",
                "publishTime": "2009-10-25T18:56:16Z"
            }
        },
        {
            "kind": "youtube#searchResult",
            "etag": "2yvIHKiDulJevuy-reVyjeRa4uE",
            "id": {
                "kind": "youtube#video",
                "videoId": "OCnXJ8O7OOo"
            },
            "snippet": {
                "publishedAt": "2021-05-23T06:39:41Z",
                "channelId": "UCH3RIG-T1TLqKOf2BNfU-ZQ",
                "title": "[เเปลไทย/SubThai] Rick Astley - Never Gonna Give You Up",
                "description": "rickastley #rickroll #subthai #เเปลไทย #ซับไทย #thaitranslate #nevergonnagiveuup #nevergonnagiveyouup #doopies #translate.",
                "thumbnails": {
                    "default": {
                        "url": "https://i.ytimg.com/vi/OCnXJ8O7OOo/default.jpg",
                        "width": 120,
                        "height": 90
                    },
                    "medium": {
                        "url": "https://i.ytimg.com/vi/OCnXJ8O7OOo/mqdefault.jpg",
                        "width": 320,
                        "height": 180
                    },
                    "high": {
                        "url": "https://i.ytimg.com/vi/OCnXJ8O7OOo/hqdefault.jpg",
                        "width": 480,
                        "height": 360
                    }
                },
                "channelTitle": "Doopies",
                "liveBroadcastContent": "none",
                "publishTime": "2021-05-23T06:39:41Z"
            }
        }
    ]`

var VideoCategoryResponseByCategory = `{
    "kind": "youtube#videoCategoryListResponse",
    "etag": "s6RguhiCzdsBQFsS4YzslvqtQtI",
	"nextPageToken": "",
	"prevPageToken": "",
    "pageInfo": {
        "totalResults": 0,
        "resultsPerPage": 0
    },
	"regionCode": "",
    "items": [
        {
            "kind": "youtube#videoCategory",
            "etag": "grPOPYEUUZN3ltuDUGEWlrTR90U",
            "id": "1",
            "snippet": {
                "title": "Film & Animation",
                "assignable": true,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        }
    ]
}`

var VideoCategoryResponseByRegion = `{
    "kind": "youtube#videoCategoryListResponse",
    "etag": "fIIv2-q7-AkaOeJf0LPrlnu-0As",
	"nextPageToken": "",
	"prevPageToken": "",
	"pageInfo": {
        "totalResults": 0,
        "resultsPerPage": 0
    },
	"regionCode": "",
    "items": [
        {
            "kind": "youtube#videoCategory",
            "etag": "grPOPYEUUZN3ltuDUGEWlrTR90U",
            "id": "1",
            "snippet": {
                "title": "Film & Animation",
                "assignable": true,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "Q0xgUf8BFM8rW3W0R9wNq809xyA",
            "id": "2",
            "snippet": {
                "title": "Autos & Vehicles",
                "assignable": true,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "qnpwjh5QlWM5hrnZCvHisquztC4",
            "id": "10",
            "snippet": {
                "title": "Music",
                "assignable": true,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "HyFIixS5BZaoBdkQdLzPdoXWipg",
            "id": "15",
            "snippet": {
                "title": "Pets & Animals",
                "assignable": true,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "PNU8SwXhjsF90fmkilVohofOi4I",
            "id": "17",
            "snippet": {
                "title": "Sports",
                "assignable": true,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "5kFljz9YJ4lEgSfVwHWi5kTAwAs",
            "id": "18",
            "snippet": {
                "title": "Short Movies",
                "assignable": false,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "ANnLQyzEA_9m3bMyJXMhKTCOiyg",
            "id": "19",
            "snippet": {
                "title": "Travel & Events",
                "assignable": true,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "0Hh6gbZ9zWjnV3sfdZjKB5LQr6E",
            "id": "20",
            "snippet": {
                "title": "Gaming",
                "assignable": true,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "q8Cp4pUfCD8Fuh8VJ_yl5cBCVNw",
            "id": "21",
            "snippet": {
                "title": "Videoblogging",
                "assignable": false,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "cHDaaqPDZsJT1FPr1-MwtyIhR28",
            "id": "22",
            "snippet": {
                "title": "People & Blogs",
                "assignable": true,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "3Uz364xBbKY50a2s0XQlv-gXJds",
            "id": "23",
            "snippet": {
                "title": "Comedy",
                "assignable": true,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "0srcLUqQzO7-NGLF7QnhdVzJQmY",
            "id": "24",
            "snippet": {
                "title": "Entertainment",
                "assignable": true,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "bQlQMjmYX7DyFkX4w3kT0osJyIc",
            "id": "25",
            "snippet": {
                "title": "News & Politics",
                "assignable": true,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "Y06N41HP_WlZmeREZvkGF0HW5pg",
            "id": "26",
            "snippet": {
                "title": "Howto & Style",
                "assignable": true,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "yBaNkLx4sX9NcDmFgAmxQcV4Y30",
            "id": "27",
            "snippet": {
                "title": "Education",
                "assignable": true,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "Mxy3A-SkmnR7MhJDZRS4DuAIbQA",
            "id": "28",
            "snippet": {
                "title": "Science & Technology",
                "assignable": true,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "4pIHL_AdN2kO7btAGAP1TvPucNk",
            "id": "30",
            "snippet": {
                "title": "Movies",
                "assignable": false,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "Iqol1myDwh2AuOnxjtn2AfYwJTU",
            "id": "31",
            "snippet": {
                "title": "Anime/Animation",
                "assignable": false,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "tzhBKCBcYWZLPai5INY4id91ss8",
            "id": "32",
            "snippet": {
                "title": "Action/Adventure",
                "assignable": false,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "ii8nBGYpKyl6FyzP3cmBCevdrbs",
            "id": "33",
            "snippet": {
                "title": "Classics",
                "assignable": false,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "Y0u9UAQCCGp60G11Arac5Mp46z4",
            "id": "34",
            "snippet": {
                "title": "Comedy",
                "assignable": false,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "_YDnyT205AMuX8etu8loOiQjbD4",
            "id": "35",
            "snippet": {
                "title": "Documentary",
                "assignable": false,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "eAl2b-uqIGRDgnlMa0EsGZjXmWg",
            "id": "36",
            "snippet": {
                "title": "Drama",
                "assignable": false,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "HDAW2HFOt3SqeDI00X-eL7OELfY",
            "id": "37",
            "snippet": {
                "title": "Family",
                "assignable": false,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "QHiWh3niw5hjDrim85M8IGF45eE",
            "id": "38",
            "snippet": {
                "title": "Foreign",
                "assignable": false,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "ztKcSS7GpH9uEyZk9nQCdNujvGg",
            "id": "39",
            "snippet": {
                "title": "Horror",
                "assignable": false,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "Ids1sm8QFeSo_cDlpcUNrnEBYWA",
            "id": "40",
            "snippet": {
                "title": "Sci-Fi/Fantasy",
                "assignable": false,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "qhfgS7MzzZHIy_UZ1dlawl1GbnY",
            "id": "41",
            "snippet": {
                "title": "Thriller",
                "assignable": false,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "TxVSfGoUyT7CJ7h7ebjg4vhIt6g",
            "id": "42",
            "snippet": {
                "title": "Shorts",
                "assignable": false,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "o9w6eNqzjHPnNbKDujnQd8pklXM",
            "id": "43",
            "snippet": {
                "title": "Shows",
                "assignable": false,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        },
        {
            "kind": "youtube#videoCategory",
            "etag": "mLdyKd0VgXKDI6GevTLBAcvRlIU",
            "id": "44",
            "snippet": {
                "title": "Trailers",
                "assignable": false,
                "channelId": "UCBR8-60-B28hp2BmDPdntcQ"
            }
        }
    ]
}`

var ChannelResponse = fmt.Sprintf(`{
    "kind": "youtube#channelListResponse",
    "etag": "00rzQGYesf5QluclbFTC8tr10MA",
    "pageInfo": {
        "totalResults": 1,
        "resultsPerPage": 5
    },
	"nextPageToken": "",
	"prevPageToken": "",
	"regionCode": "",
	"items": %s
}`, ChannelResponseItems)

var ChannelResponseItems = `[
        {
            "kind": "youtube#channel",
            "etag": "5N2NXrPKn7QMietcmPOavHlHoFI",
            "id": "UCuAXFkgsw1L7xaCfnd5JJOw",
            "snippet": {
                "title": "Rick Astley",
                "description": "Official YouTube channel for Rick Astley.",
                "customUrl": "@rickastleycoukofficial",
                "publishedAt": "2015-02-01T16:32:30Z",
                "thumbnails": {
                    "default": {
                        "url": "https://yt3.ggpht.com/BbWaWU-qyR5nfxxXclxsI8zepppYL5x1agIPGfRdXFm5fPEewDsRRWg4x6P6fdKNhj84GoUpUI4=s88-c-k-c0x00ffffff-no-nd-rj",
                        "width": 88,
                        "height": 88
                    },
                    "medium": {
                        "url": "https://yt3.ggpht.com/BbWaWU-qyR5nfxxXclxsI8zepppYL5x1agIPGfRdXFm5fPEewDsRRWg4x6P6fdKNhj84GoUpUI4=s240-c-k-c0x00ffffff-no-nd-rj",
                        "width": 240,
                        "height": 240
                    },
                    "high": {
                        "url": "https://yt3.ggpht.com/BbWaWU-qyR5nfxxXclxsI8zepppYL5x1agIPGfRdXFm5fPEewDsRRWg4x6P6fdKNhj84GoUpUI4=s800-c-k-c0x00ffffff-no-nd-rj",
                        "width": 800,
                        "height": 800
                    }
                },
                "localized": {
                    "title": "Rick Astley",
                    "description": "Official YouTube channel for Rick Astley."
                },
                "country": "GB"
            },
            "contentDetails": {
                "relatedPlaylists": {
                    "likes": "",
                    "uploads": "UUuAXFkgsw1L7xaCfnd5JJOw"
                }
            },
            "statistics": {
                "viewCount": "1675408356",
                "subscriberCount": "3430000",
                "hiddenSubscriberCount": false,
                "videoCount": "91"
            },
            "topicDetails": {
                "topicIds": [
                    "/m/04rlf",
                    "/m/064t9",
                    "/m/02mscn"
                ],
                "topicCategories": [
                    "https://en.wikipedia.org/wiki/Music",
                    "https://en.wikipedia.org/wiki/Pop_music",
                    "https://en.wikipedia.org/wiki/Christian_music"
                ]
            },
            "status": {
                "privacyStatus": "public",
                "isLinked": true,
                "longUploadsStatus": "longUploadsUnspecified",
                "madeForKids": false
            },
            "brandingSettings": {
                "channel": {
                    "title": "Rick Astley",
                    "description": "Official YouTube channel for Rick Astley.",
                    "keywords": "Official Rick Astley \"rick astley\" \"never gonna give you up\" \"the best of me\" \"best of me\" \"together forever\" \"cry for help\" \"beautiful life\" \"official video\" \"official audio\" rickroll rickrolled \"rick roll\" \"rick asltey\" \"give up\" meme",
                    "unsubscribedTrailer": "m1k3Cpke4yU",
                    "country": "GB",
					"trackingAnalyticsAccountId": ""
                },
                "image": {
                    "bannerExternalUrl": "https://yt3.ggpht.com/svWR3D0zc85G3GEhuHVVY2Bj5zHElnhi8Sw_CKwSfInBddUdElbytQt-UYOVb2XR0QG03aAsKzU"
                }
            }
        }
    ]`
