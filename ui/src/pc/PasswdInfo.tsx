import React from 'React'
import {Button, Table} from "antd";
import BlankElement from "./component/BlankElement";

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

    render() {
        return (<div>
            <Table
                className="components-table-demo-nested"
                columns={this.ui_table.columns}
                dataSource={this.state.tableData}/>
        </div>);
    }

}