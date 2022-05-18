import axios from "axios";
import ObjectUtil from "../utils/ObjectUtil";
import {notification} from "antd";
import Resp from "../model/HttpModel";
import RouterURL from "../env/RouterURL";
import RouterUtil from "../utils/RouterUtil";

axios.interceptors.response.use(
    (response: any) => {
        const pathName = RouterUtil.getPath();
        if (!response.data.login && pathName != RouterURL.LOGIN) {
            RouterUtil.push(RouterURL.LOGIN);
        } else {
            return response;
        }
    }
);

export default class HttpClient {

    private static successRespFunc = (url: string, value: any, successFunc?: (r: Resp) => void, finalFunc?: Function) => {
        if (ObjectUtil.isNotNull(value)) {
            let _data = value.data;
            if (_data.status === 'error') {
                console.log('errLog[' + url + ']:' + _data.data.errormsg);
                notification.open({
                    message: '错误信息',
                    description: _data.data.errormsg,
                    duration: 1.5
                });
            } else {
                if (successFunc) {
                    successFunc(_data as Resp);
                }
            }
        }
        if (finalFunc) {
            finalFunc();
        }
    }

    private static errorRespFunc = (err: any, finalFunc?: (err: any) => void) => {
        console.log(err);
        if (finalFunc) {
            finalFunc(err);
        }
    }

    public static post = (url: string, data: {} = {}, successFunc?: (resp: Resp) => void, finalFunc?: (err?: any) => void): void => {
        axios.post(url, data).then(value => {
            HttpClient.successRespFunc(url, value, successFunc, finalFunc)
        }).catch(err => {
            HttpClient.errorRespFunc(err, finalFunc)
        });
    }

    public static get = (url: string, successFunc?: (resp: Resp) => void, finalFunc?: (err?: any) => void): void => {
        axios.get(url).then(value => {
            HttpClient.successRespFunc(url, value, successFunc, finalFunc)
        }).catch(err => {
            HttpClient.errorRespFunc(err, finalFunc)
        });
    }
}