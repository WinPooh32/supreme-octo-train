import * as React from "react";
import { Menu } from "semantic-ui-react";
import { Redirect } from "react-router";

export interface MenuState { activeItem: string, redirect: boolean }

type Indexable = { [label: string]: string };
const links = {
    'home': '/page/1',
    'about': '/about/'
} as Indexable

export class MyMenu extends React.Component<any, MenuState> {
    handleItemClick = (e: any, { name }: any) => {
        this.setState({ activeItem: name, redirect: true } as MenuState)
    }

    constructor(props: any) {
        super(props, {})
        this.state = { activeItem: "editorials" } as MenuState
    }

    render() {
        const { activeItem, redirect } = this.state

        if (redirect) {
            this.setState({ activeItem: name, redirect: false } as MenuState)
            return <Redirect to={`${links[activeItem]}`} />
        }

        return (
            <Menu>
                <Menu.Item
                    name='home'
                    active={activeItem === 'home'}
                    onClick={this.handleItemClick}
                >
                    Товары
                </Menu.Item>
                {/* 
                <Menu.Item name='about' active={activeItem === 'about'} onClick={this.handleItemClick}>
                    Статистика
                </Menu.Item>

                <Menu.Item
                    name='upcomingEvents'
                    active={activeItem === 'upcomingEvents'}
                    onClick={this.handleItemClick}
                >
                    Upcoming Events
                </Menu.Item> */}
            </Menu>
        )
    }
}