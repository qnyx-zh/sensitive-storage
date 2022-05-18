import {Input} from "antd";
import React from "react";
import {InputProps} from "antd/lib/input/Input";



export default class InputEx extends React.Component<InputProps & {
    rtModel : any
}, any> {

    constructor(props: any) {
        super(props);
    }

    render() {
        return (
            <Input onChange={(val) => {
                const value = val.target.value;
            }}/>
        )
    }
}