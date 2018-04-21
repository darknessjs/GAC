import { Button, Input } from 'antd';
import * as React from 'react';


class Result extends React.Component {
    public render() {
        return (
            <div className="result-main">
                <div className="result-main-cell">请输入用户id</div>
                <Input />
                <div className="result-main-cell">请输入用户名</div>
                <Input />
                <Button> 确定</Button>
            </div>
        );
    }
}

export default Result;
