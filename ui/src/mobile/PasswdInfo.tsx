import React from 'react';
import {Button, SearchBar, Space} from "antd-mobile";
import css from './style/passwd-info.module.css';

class State {

}

export class PasswdInfo extends React.Component<any, State> {

    constructor(props: any) {
        super(props);
        this.state = new State()
    }

    ui_list = {
        render: () => {
            const elements = []
            for (let i = 0; i < 10; i++) {
                elements.push(<div className={css.passwdLine}>
                    <div className={css.topic}>这是一个密码信息标题</div>
                    <div className={css.date}>
                        <Space block justify={"between"}>
                            <div>kiririx@github.com</div>
                            <div>2022-03-23 12:99:99</div>
                        </Space>
                    </div>
                </div>)
            }
            return elements
        }
    }

    render() {
        return <div className={css.passwdBody}>
            <div>
                <SearchBar placeholder='请输入标题' style={{ '--background': '#ffffff' }}/>
            </div>
            <div>
                {this.ui_list.render()}
            </div>
        </div>
    }

}
