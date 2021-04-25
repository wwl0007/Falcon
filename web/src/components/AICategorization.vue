<template>
    <div class="d-flex justify-content-center">
        <div class="d-flex flex-column">
            <p class="fs-4 text-center align-items-center">Cancer Risk</p>

            <div class="slide-container d-flex justify-content-center mb-2">
                <span>0%</span>
                <input
                    class="my-slider ms-2 me-2"
                    type="range"
                    min="0"
                    max="100"
                    :value="value"
                    @input="onInput"
                    step="33.33333"
                />
                <span>100%</span>
            </div>

<!--            <p class="text-muted text-center">Confidence: {{ confidence }}%</p>-->
            <p class="text-muted fs-6 text-center">Tip: drag the slider to give the AI feedback.</p>
        </div>
    </div>
</template>

<script lang="ts">
    import { Component, Prop, Vue } from "vue-property-decorator";

    @Component
    export default class AICategorization extends Vue {
        @Prop() confidence!: number;
        @Prop() value!: number;

        onInput(event: any) {
            this.$emit('input', parseInt(event.target.value));
        }
    }
</script>

<style scoped lang="scss">
    .my-slider {

        width: 35vw;
    }

    /*Chrome*/
    @media screen and (-webkit-min-device-pixel-ratio:0) {
        input[type='range'] {
            overflow: hidden;
            width: 35vw;
            -webkit-appearance: none;
            background: linear-gradient(90deg, rgb(0 255 0) 0%, rgba(255,0,0,1) 100%);
        }

        input[type='range']::-webkit-slider-runnable-track {
            height: 10px;
            -webkit-appearance: none;
            color: #13bba4;
            margin-top: -1px;
        }

        input[type='range']::-webkit-slider-thumb {
            width: 10px;
            -webkit-appearance: none;
            margin-top: -20px;
            height: 100px;
            cursor: ew-resize;
            background: #434343;
            //box-shadow: -80px 0 0 80px #43e5f7;
        }

    }
</style>
