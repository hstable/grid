import axios from "@/plugins/axios";
import { handleResponse } from "./utils";
import { getUserInfoFromToken } from "@/assets/js/tokenTools";

function login(Username, Password) {
  return axios({
    url: apiRoot + "/user/login",
    method: "post",
    data: {
      Username,
      Password
    }
  }).then(res => {
    handleResponse(res, this, () => {
      localStorage["token"] = res.data.data.Token;
      let o = getUserInfoFromToken(res.data.data.Token);
      localStorage["name"] = o.name;
      localStorage["sub"] = o.sub;
    });
    return res;
  });
}

async function nextImage(After) {
  let data = null;
  await axios({
    url: apiRoot + "/image/next",
    method: "get",
    params: {
      After
    }
  }).then(res => {
    handleResponse(res, this, () => {
      data = res.data.data;
    });
  });
  return data;
}
async function lastImage(Before) {
  let data = null;
  await axios({
    url: apiRoot + "/image/last",
    method: "get",
    params: {
      Before
    }
  }).then(res => {
    handleResponse(res, this, () => {
      data = res.data.data;
    });
  });
  return data;
}

function mark(ID, Mark) {
  return axios({
    url: apiRoot + "/image/mark",
    method: "put",
    data: {
      ID,
      Mark
    }
  });
}

function putStandard(standard) {
  return axios({
    url: apiRoot + "/user/standard",
    method: "put",
    data: {
      standard: JSON.stringify(standard)
    }
  });
}
async function getOwnStandard() {
  let standard = null;
  await axios({
    url: apiRoot + "/user/standard/own",
    method: "get"
  }).then(res => {
    handleResponse(res, this, () => {
      standard = JSON.parse(res.data.data.Standard);
    },()=>{});
    return res;
  });
  return standard;
}

export default {
  login,
  nextImage,
  lastImage,
  mark,
  putStandard,
  getOwnStandard
};
