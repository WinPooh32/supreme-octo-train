import * as React from "react";
import { LineChart, CartesianGrid, XAxis, YAxis, Tooltip, Legend, Line, AreaChart, Brush, ReferenceLine } from "recharts";
import { Form, InputOnChangeData, Grid, Statistic, Divider, List, Icon } from "semantic-ui-react";
import { Link } from "react-router-dom";


type Indexable = { [label: string]: any };


function makePlotData(item: any): Array<Indexable> {
    let size = item.data.length
    let data = item.data
    let forecast = item.forecast
    let filtered = item.filtered
    let restored = item.restored

    let pts = new Array<Indexable>(size)

    for (let i = 0; i < size; ++i) {
        let pt = {
            name: `${Math.floor(i / 52) + 1}-${Math.floor(i % 52) + 1}`,
            'Продажи': data[i],
            'Сглаж.': (filtered[i] > 0 ? Math.ceil(filtered[i]) : 0),
            'Восст.': (restored[i] > 0 ? restored[i] : null),
        } as Indexable;

        if (i >= size - 52) {
            pt['Прогноз'] = Math.ceil(forecast[i % 52])
        }

        // console.log(pt)

        pts[i] = pt
    }

    return pts
}

//Hack
var instance: ViewItemInfo

export interface ViewItemInfoProps { item: any; }
export interface ViewItemInfoState {
    input?: Indexable,
    sum: number,
    sumSold: number,
    sumError: number,
    isNextItem: boolean,
    prevItem: any,
    data: Indexable[]
}

export class ViewItemInfo extends React.Component<ViewItemInfoProps, ViewItemInfoState> {
    startIndex: number
    endIndex: number
    private onChangeItem(e: React.SyntheticEvent) {
        this.setState((state) => {
            return { isNextItem: true }
        })
    }

    private onBrushChange(...args: any[]) {
        // console.log(args[0].startIndex, args[0].endIndex)
        // instance.setState((state) => {
        //     return { startIndex: args[0].startIndex, endIndex: args[0].endIndex } as ViewItemInfoState
        // })
        instance.startIndex = args[0].startIndex
        instance.endIndex = args[0].endIndex
    }

    private calcSum(data: Indexable[], begin: number, end: number): number {
        let sum = 0

        for (let i = begin; i < end; ++i) {
            sum += data[52 * 3 + i]['Прогноз'] as number
        }

        return sum
    }


    private calcSoldSum(data: Indexable[], begin: number, end: number): number {
        let sum = 0

        for (let i = begin; i < end; ++i) {
            sum += data[52 * 3 + i]['Продажи'] as number
        }

        return sum
    }

    private onCalcInput(event: React.ChangeEvent<HTMLInputElement>, data: InputOnChangeData) {
        instance.setState((state) => {
            let b = parseInt(state.input.calcBegin as string)
            let e = parseInt(state.input.calcEnd as string)

            let sum = instance.calcSum(state.data, b, e)
            let sumSold = instance.calcSoldSum(state.data, b, e)
            let sumError: number

            if (sumSold != 0) {
                sumError = 100 * (sum - sumSold) / sumSold
            } else {
                sumError = Infinity
            }

            let newState = {
                input: state.input,
                sum: sum,
                sumSold: sumSold,
                sumError: sumError,
                data: state.data
            }

            newState.input[data.name] = data.value

            return newState
        })
    }

    constructor(props: any) {
        super(props)

        // this.onBrushChange.bind(this)
        // this.onCalcInput.bind(this)
        // this.onChangeItem.bind(this)

        this.startIndex = 0
        this.endIndex = 52 * 4 - 1

        let item = props.item

        let data = makePlotData(item)

        let sum = this.calcSum(data, 1, 52)
        let sumSold = this.calcSoldSum(data, 1, 52)
        let sumError = 100 * (sum - sumSold) / sumSold

        this.state = {
            input: { calcBegin: '1', calcEnd: '52' } as Indexable,
            sum: sum,
            sumSold: sumSold,
            sumError: sumError,
            prevItem: item,
            data: data
        } as ViewItemInfoState

        instance = this
    }

    componentDidUpdate() {
        if (this.props.item != this.state.prevItem) {
            this.setState((state) => {
                let item = this.props.item
                let data = makePlotData(item)
                let sum = this.calcSum(data, 1, 52)
                let sumSold = this.calcSoldSum(data, 1, 52)
                let sumError = 100 * (sum - sumSold) / sumSold

                return {
                    input: { calcBegin: '1', calcEnd: '52' } as Indexable,
                    sum: sum,
                    sumSold: sumSold,
                    sumError: sumError,
                    isNextItem: this.props.item,
                    data: data,
                    prevItem: item
                } as ViewItemInfoState
            })
        }
    }

    render() {
        let item = this.props.item

        let width = 1127
        let height = 400

        return (
            <span>
                <List horizontal={true}>
                    <List.Item><Link className="ui blue button" to={`../item/${item.id}`}> <Icon name='arrow left' size='large' /> </Link></List.Item>
                    <List.Item><Link className="ui blue button" to={`../item/${item.id + 2}`}> <Icon name='arrow right' size='large' /> </Link></List.Item>
                </List>

                <h1> {item.name} </h1>

                <br />

                <Grid columns={2}>
                    <Grid.Column as={Form}>
                        <Form.Input
                            label={`Начало: ${this.state.input.calcBegin}`}
                            min={1}
                            max={52}
                            name='calcBegin'
                            onChange={this.onCalcInput}
                            step={1}
                            type='range'
                            value={this.state.input.calcBegin}
                        />

                        <Form.Input
                            label={`Конец: ${this.state.input.calcEnd}`}
                            min={1}
                            max={52}
                            name='calcEnd'
                            onChange={this.onCalcInput}
                            step={1}
                            type='range'
                            value={this.state.input.calcEnd}
                        />

                        <Statistic size='mini'>
                            <Statistic.Label>Заказано:</Statistic.Label>
                            <Statistic.Value>{Math.ceil(this.state.sum)}</Statistic.Value>
                        </Statistic>

                        <Statistic size='mini'>
                            <Statistic.Label>Продано:</Statistic.Label>
                            <Statistic.Value>{Math.ceil(this.state.sumSold)}</Statistic.Value>
                        </Statistic>

                        <Statistic size='mini'>
                            <Statistic.Label>Ошибка:</Statistic.Label>
                            <Statistic.Value>{Math.ceil(this.state.sumError)}%</Statistic.Value>
                        </Statistic>

                        <br />

                    </Grid.Column>
                </Grid>

                <LineChart
                    width={width}
                    height={height}
                    data={this.state.data}
                    margin={{
                        top: 5, right: 30, left: 20, bottom: 5,
                    }}
                >
                    <CartesianGrid strokeDasharray="3 3" />
                    <XAxis dataKey="name" />
                    <YAxis />
                    <Tooltip />
                    <Legend />

                    <ReferenceLine x="2-1" stroke="red" />
                    <ReferenceLine x="3-1" stroke="red" />
                    <ReferenceLine x="4-1" stroke="red" />

                    <ReferenceLine x={`4-${this.state.input.calcBegin}`} stroke="magenta" />
                    <ReferenceLine x={`4-${this.state.input.calcEnd}`} stroke="magenta" />

                    <Line type="monotone" dataKey="Продажи" stroke="#8884d8" activeDot={{ r: 8 }} />
                    <Line type="monotone" dataKey="Сглаж." stroke="#82ca9d" />
                    <Line type="monotone" dataKey="Прогноз" stroke="#ffc658" />
                    <Line type="monotone" dataKey="Восст." stroke="#ff0000" />

                    <Brush startIndex={this.startIndex} endIndex={this.endIndex} onChange={this.onBrushChange} />
                </LineChart>

            </span>
        )
    }
}