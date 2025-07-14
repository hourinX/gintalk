import {WechatOutlined,GithubOutlined} from '@ant-design/icons-vue'
import { install } from 'ant-design-vue'

const icons = [
    WechatOutlined,GithubOutlined
]

export default {
    install(app) {
        icons.forEach((icon) => {
            app.component(icon.displayName, icon);
        })
    }
}