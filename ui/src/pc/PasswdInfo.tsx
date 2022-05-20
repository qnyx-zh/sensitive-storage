import React from 'react'
import {Button, Form, Input, Modal, notification, Popconfirm, Row, Table, Tooltip} from "antd";
import BlankElement from "./component/BlankElement";
import BlankRow from "./component/BlankRow";
import HttpClient from "../common/component/HttpClient";
import Search from "antd/es/input/Search";
import {CopyOutlined, LockOutlined, UnlockOutlined} from '@ant-design/icons';
import css from './style/passwd-info.module.css';
import ObjectUtil from "../common/utils/ObjectUtil";
import Resp from "../common/model/HttpModel";
import HttpURL from "../common/env/HttpURL";

class TableRow {
    id = ""
    username = ""
    password = ""
    topic = ""
    description = ""
    locked = true
}

class FormData {
    username = ""
    password = ""
    topic = ""
    description = ""
}

class State {
    tableData: Array<TableRow> = []
    dialogVisible: boolean = false
    formData: FormData = new FormData()
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
                render: (col: any, row: any, index: any) => {
                    return <div>
                        <label>{row.topic}</label>
                    </div>
                }
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
                render: (col: any, row: TableRow, index: number) => {
                    return <div>
                        <span>
                          <label>{row.locked ? row['password'] : '*******'}</label>
                        </span>
                        <span className={css.passwdLock} onClick={() => {
                            const _tableData = [...this.state.tableData]
                            for (const i in this.state.tableData) {
                                if (parseInt(i) === index) {
                                    const locked: boolean = _tableData[i].locked
                                    _tableData[i] = {
                                        ..._tableData[i],
                                        locked: !locked
                                    }
                                    this.setState({
                                        tableData: _tableData
                                    })
                                    break
                                }
                            }
                        }
                        }>
                            {row.locked ? <UnlockOutlined/> : <LockOutlined/>}
                        </span>
                        <Tooltip placement="top" title={"复制密码"}>
                            <span className={css.passwdCopy} onClick={(v) => {
                                const inp = document.createElement("input")
                                const val: string = row['password']
                                inp.setAttribute('value', val)
                                inp.setAttribute('style', 'display:\'none\'')
                                document.body.append(inp)
                                inp.select()
                                document.execCommand("copy")
                                document.body.removeChild(inp)
                                notification.open({
                                    message: '提示',
                                    description: '密码已复制到剪切板',
                                    duration: 1.5
                                });
                            }
                            }>
                             <CopyOutlined/>
                        </span>
                        </Tooltip>
                    </div>
                }
            },
            {
                title: '描述',
                dataIndex: 'description',
                key: 'description',
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
                        <Popconfirm placement="topLeft" title={"确定删除吗？"} onConfirm={() => {
                            this.ui_table.doDelete(row.id)
                        }
                        } okText="是" cancelText="否">
                            <Button onClick={() => {
                            }}>删除</Button>
                        </Popconfirm>
                    </div>;
                }
            }
        ],
        doDelete: (id: string) => {
            const url = HttpURL.DEL_PASSWD + "/" + id
            HttpClient.delete(url, resp => {
                this.ui_table.loadData()
            })
        },
        loadData: () => {
            const url = "http://127.0.0.1:4523/mock/991824/passwdInfos"
            HttpClient.get(url, resp => {
                let _tableData = resp.data.passwds
                _tableData = _tableData.map((v: TableRow) => {
                    return {
                        ...v,
                        locked: false
                    }
                })
                this.setState({
                    tableData: _tableData
                })
            })
        },
        doSearch: (text: string) => {
            const url = "http://127.0.0.1:4523/mock/991824/search?q=" + text
            HttpClient.get(url, resp => {
                this.setState({
                    tableData: resp.data.passwds
                })
            })
        }
    }

    protected ui_dialogs = {
        layout: {
            labelCol: {span: 5},
            wrapperCol: {span: 19},
        },
        doClose: () => {
            this.setState({
                dialogVisible: false
            })
        },
        doOk: () => {
            const url = ""
            HttpClient.post(url, this.state.formData, (resp: Resp) => {
                this.ui_dialogs.doClose();
                this.ui_table.loadData();
            });
        },
        doClear: () => {
            this.setState({
                formData: {
                    username: "",
                    password: "",
                    topic: "",
                    description: ""
                }
            })
        },
        openEditForm: (id?: string) => {
            this.ui_dialogs.doClear()
            const showFn = () => {
                this.setState({
                        dialogVisible: true
                    }
                );
            }
            if (id) {
                const url = "http://127.0.0.1:4523/mock/991824/passwdInfo/6ba7b810-9dad-11d1-80b4-00c04fd430c8"
                HttpClient.get(url, resp => {
                    this.setState({
                            formData: (resp.data as FormData)
                        }
                    )
                    showFn()
                })
            } else {
                showFn();
            }
        },
        render: () => {
            return (
                <Modal title={"编辑密码信息"}
                       onCancel={this.ui_dialogs.doClose}
                       onOk={this.ui_dialogs.doOk}
                       visible={this.state.dialogVisible}>
                    <Form {...this.ui_dialogs.layout}>
                        <Form.Item label={"标题"}>
                            <Input value={this.state.formData.topic} onChange={(val) => {
                                this.setState({formData: ObjectUtil.getNewProperty(this.state, "formData.topic", val.target.value)})
                            }}/>
                        </Form.Item>
                        <Form.Item label={"用户名"}>
                            <Input value={this.state.formData.username} onChange={(val) => {
                                this.setState({formData: ObjectUtil.getNewProperty(this.state, "formData.username", val.target.value)})
                            }}/>
                        </Form.Item>
                        <Form.Item label={"密码"}>
                            <Input value={this.state.formData.password} onChange={(val) => {
                                this.setState({formData: ObjectUtil.getNewProperty(this.state, "formData.password", val.target.value)})
                            }}/>
                        </Form.Item>
                        <Form.Item label={"描述"}>
                            <Input.TextArea value={this.state.formData.description} onChange={(val) => {
                                this.setState({formData: ObjectUtil.getNewProperty(this.state, "formData.description", val.target.value)})
                            }}/>
                        </Form.Item>
                    </Form>
                </Modal>
            )
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
                <Button type={"primary"} onClick={() => {
                    this.ui_dialogs.openEditForm()
                }}>添加</Button>
                <BlankElement/>
                <Search placeholder="请输入标题" onSearch={this.ui_table.doSearch} style={{width: 200}}/>
                <BlankElement/>
            </Row>
            <BlankRow/>
            <Table
                rowKey={"id"}
                className="components-table-demo-nested"
                columns={this.ui_table.columns}
                dataSource={this.state.tableData}/>
            {this.ui_dialogs.render()}
        </div>);
    }

}