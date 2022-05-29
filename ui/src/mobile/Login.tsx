import React from 'react';
import {Grid, Input} from "antd-mobile";

class State {
}

export class Login extends React.Component<any, State> {

    constructor(props: any) {
        super(props);
        this.state = new State()
    }

    render() {
        return <div style={{
            backgroundColor: "#f7f7f7"
        }}>
            <div style={{

            }}>
                <div style={{
                    padding: '4rem 2rem',
                    margin: '18rem 1rem',
                    borderRadius: '2rem',
                    borderColor: 'black',
                    border: 'solid',
                    backgroundColor: '#FFFFFF'
                }}>
                    <Grid columns={3} gap={2} style={{
                        margin: '1rem 0'
                    }}>
                        <Grid.Item>
                            <label>用户名: </label>
                        </Grid.Item>
                        <Grid.Item span={2}>
                            <Input placeholder={"请输入用户名"}/>
                        </Grid.Item>
                    </Grid>
                    <Grid columns={3} gap={2}>
                        <Grid.Item>
                            <label>密码: </label>
                        </Grid.Item>
                        <Grid.Item span={2}>
                            <Input placeholder={"请输入密码"}/>
                        </Grid.Item>
                    </Grid>

                </div>
            </div>
        </div>;
    }

}