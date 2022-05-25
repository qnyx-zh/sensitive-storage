import {notification} from "antd";


export class Notification {

    public static error = (message: string) => {
        notification.open({
            message: '错误信息',
            description: message,
            duration: 1.5
        });
    }

    public static success = (message: string) => {
        notification.open({
            message: '提示',
            description: message,
            duration: 1.5
        })
    }
}