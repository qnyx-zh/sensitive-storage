import React from 'react'
import {Button, Row, Table} from "antd";
import BlankElement from "./component/BlankElement";
import BlankRow from "./component/BlankRow";
import HttpClient from "../common/component/HttpClient";
import InputEx from "./component/InputEx";
import Search from "antd/es/input/Search";

class State {
    tableData: Array<{ key: string, label: string }> = []
}

class Props {
}

export class PasswdInfo extends React.Component<Props, State> {

    private ui_table = {
        columns: [
            {
                title: '标题',
                dataIndex: 'topic',
                key: 'topic',
            },
            {
                title: '用户名',
                dataIndex: 'username',
                key: 'username',
            },
            {
                title: '密码',
                dataIndex: 'password',
                key: 'password',
            },
            {
                title: '操作',
                dataIndex: 'ope',
                key: 'ope',
                render: (col: any, row: any, index: any) => {
                    return <div>
                        <Button onClick={() => {
                            this.ui_dialogs.openEditForm(row.id)
                        }}>编辑</Button>
                        <BlankElement/>
                        <Button onClick={() => {
                            this.ui_table.doDelete(row.id)
                        }}>删除</Button>
                    </div>;
                }
            }
        ],
        doDelete: (id: string) => {
            // todo
        },
        loadData: () => {
            const url = "http://127.0.0.1:4523/mock/991824/passwdInfos"
            HttpClient.get(url, resp => {
                this.setState({
                    tableData: resp.data.passwds
                })
            })
        }
    }

    private ui_dialogs = {
        openEditForm: (id: string) => {
            // 打开编辑框
        }
    }

    constructor(props: Props) {
        super(props);
        this.state = new State()
    }



    componentDidMount() {
        this.ui_table.loadData()
    }

    render() {
        return (<div>
            <Row>
                <Button type={"primary"} >添加</Button>
                <BlankElement/>
                <Button type={"primary"} >刷新</Button>
                <BlankElement/>
                <Search placeholder="input search text" onSearch={()=>{

                }} style={{ width: 200 }} />
                <BlankElement/>
            </Row>
            <BlankRow/>
            <Table
                className="components-table-demo-nested"
                columns={this.ui_table.columns}
                dataSource={this.state.tableData}/>
        </div>);
    }

}