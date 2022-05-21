import React from 'react'
import {Button, Col, Form, Input, Modal, notification, Popconfirm, Row, Table, Tooltip} from "antd";
import BlankElement from "./component/BlankElement";
import BlankRow from "./component/BlankRow";
import HttpClient from "../common/component/HttpClient";
import Search from "antd/es/input/Search";
import {CopyOutlined, LockOutlined, UnlockOutlined} from '@ant-design/icons';
import css from './style/passwd-info.module.css';
import ObjectUtil from "../common/utils/ObjectUtil";
import Resp from "../common/model/HttpModel";
import HttpURL from "../common/env/HttpURL";
import StringUtil from "../common/utils/StringUtil";
import {Notification} from "../common/component/Notification";

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
                width: 250,
                render: (col: any, row: TableRow, index: number) => {
                    return <div>
                        <Row>
                            <Col span={16}>
                                <span>
                          <label>{row.locked ? row['password'] : '**************'}</label>
                        </span>
                            </Col>
                            <Col span={8}>
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
                            </Col>
                        </Row>
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
            const url = HttpURL.GET_PASSWD_LIST
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
            const url = HttpURL.GET_PASSWD_SEARCH + "?q=" + text
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
        // 对必填字段进行验证，通过返回true，未通过返回false
        doValidate: () => {
            const validateFields = ["username", "password", "topic"]
            let key: keyof FormData
            for (key in this.state.formData) {
                if (validateFields.indexOf(key) > -1 && StringUtil.isBlank(this.state.formData[key])) {
                    return false
                }
            }
            return true
        },
        doOk: () => {
            if (!this.ui_dialogs.doValidate()) {
                Notification.error("请填写完整信息")
                return
            }
            const url = HttpURL.POST_PASSWD
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
                const url = HttpURL.GET_PASSWD + "/" + id
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
                        <Form.Item label={"标题"} name={"topic"} rules={[{required: true, message: '请输入标题'}]}>
                            <Input value={this.state.formData.topic} onChange={(val) => {
                                this.setState({formData: ObjectUtil.getNewProperty(this.state, "formData.topic", val.target.value)})
                            }}/>
                        </Form.Item>
                        <Form.Item label={"用户名"} name={"username"} rules={[{required: true, message: '请输入用户名'}]}>
                            <Input value={this.state.formData.username} onChange={(val) => {
                                this.setState({formData: ObjectUtil.getNewProperty(this.state, "formData.username", val.target.value)})
                            }}/>
                        </Form.Item>
                        <Form.Item label={"密码"} name={"password"} rules={[{required: true, message: '请输入密码'}]}>
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
        return (<div className={css.passwdBody}>
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