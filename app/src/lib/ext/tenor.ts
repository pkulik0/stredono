import {PUBLIC_TENOR_KEY} from "$env/static/public";
import axios from "axios";

export interface Gif {
    id: string;
    title: string;
    url: string;
    preview: string;
}

const tenorUrl = "https://g.tenor.com/v1"

const tenorResponseToGifs = (data: any): Gif[] => {
    return data.map((gif: any) => {
        return {
            id: gif.id,
            title: gif.title,
            url: gif.media[0].gif.url,
            preview: gif.media[0].tinygif.url
        }
    });

}

const getTenorUrl = (endpoint: string, limit: number, next: string) => {
    let url = `${tenorUrl}/${endpoint}?key=${PUBLIC_TENOR_KEY}&contentfilter=high&limit=${limit}&ar_range=wide`;
    if (next) {
        url += `&pos=${next}`;
    }
    return url;
}

export interface GifsResponse {
    gifs: Gif[];
    next: string;
}

export const getTopGifs = async (limit: number = 20, next: string = ""): Promise<GifsResponse> => {
    const url = getTenorUrl("trending", limit, next);
    const res = await axios.get(url);
    return {
        gifs: tenorResponseToGifs(res.data.results),
        next: res.data.next
    };
}

export const searchGifs = async (query: string, limit: number = 20, next: string = ""): Promise<GifsResponse> => {
    const url = getTenorUrl("search", limit, next);
    const res = await axios.get(`${url}&q=${query}`);
    return {
        gifs: tenorResponseToGifs(res.data.results),
        next: res.data.next
    };
}