import React, { useState, useMemo } from "react";
import api from '../services/api';
import { MailProperty, MailPattern } from "../@types/mail";
import { Modal, Layout, Form, Input, notification } from 'antd';

const { TextArea } = Input;

function Modall({show, handleAction}: any): any
{
    const [form] = Form.useForm();
    useMemo((): any => {}, [ handleAction ]);

    const handleOk = async () => {
        form.validateFields().then(async (values) => {
            await api.post("/mail", mail)
            .then((response: any) => {
                notification.open({
                    message: "Email successfully sent",
                  });
            })
            .catch( (err: any) => {
                console.log('Invalid!');
            });
        }).catch();
    };
  
    const [mail, setMail]: any = useState<MailProperty>(MailPattern);

    return(
        <>
            <Layout>
                <Modal title="Send Mail" visible={show} onOk={handleOk} onCancel={handleAction}>
                    <Form
                        name="basic"
                        labelCol={{ span: 8 }}
                        wrapperCol={{ span: 16 }}
                        initialValues={{ remember: true }}
                        onFinish={handleOk}
                        form={form}
                        autoComplete="off"
                    >
                        <Form.Item
                            name="name"
                            rules={[{ required: true, message: 'Please input your name!' }]}
                        >
                            <Input
                                placeholder="Subject"
                                type="text"
                                onChange={(val: any) => setMail({ ...mail, subject: val.target.value })}
                            />
                        </Form.Item>
                        <Form.Item
                            name="email"
                            rules={[{ required: true, message: 'Please input your email!' }]}
                        >
                            <Input
                                placeholder="Email"
                                type="email"
                                onChange={(val: any) => setMail({ ...mail, email: val.target.value })}
                            />
                        </Form.Item>
                        <Form.Item
                            name="message"
                            rules={[{ required: true, message: 'Please input your message!' }]}
                        >
                            <TextArea
                                placeholder="Message"
                                onChange={(val: any) => setMail({ ...mail, message: val.target.value })}
                                rows={4}
                            />
                        </Form.Item>
                    </Form>
                </Modal>
            </Layout>
        </>
    );
}

export default Modall;