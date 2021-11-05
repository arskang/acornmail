# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- FontSize ```acorntypes.FontSize```
- CouponParams ```acorntypes.CouponParams```
- PromoItems ```acorntypes.PromoItems```
- PromoParams ```acorntypes.PromoParams```
- CouponStyles ```acorntypes.CouponStyles```
- GetSizes
- NewAcornSize
- GetAligns
- NewCoupon
- NewPromo
- Carpeta con ejemplos
- Documentación a los métodos
- Archivo **.editorconfig**
- Archivo **CHANGELOG.md**

### Changed

- La importación cambia a github.com/arskang/acornmail
- Se optimizaron las imagenes con ejemplos en el **README.md**
- Se renombró la carpeta **assets** por **_assets** para que golang no lo importe

### Removed

- ```acorntypes.LabelStyles.Type``` usar mejor ```acorntypes.LabelStyles.Outlined```

## [1.0.0-beta] - 2021-10-23

### Added

- Package **acorntypes**
    - Color ```acorntypes.Color```
    - WidthColumn ```acorntypes.WidthColumn```
    - Types ```acorntypes.Types```
    - AcornComponents ```acorntypes.AcornComponents```
    - AcornVariables ```acorntypes.AcornVariables```
    - ColumnParams ```acorntypes.ColumnParams```
    - ColumnStyles ```acorntypes.ColumnStyles```
    - ButtonParams ```acorntypes.ButtonParams```
    - ButtonStyles ```acorntypes.ButtonStyles```
    - AlertParams ```acorntypes.AlertParams```
    - AlertStyles ```acorntypes.AlertStyles```
    - AccordionParams ```acorntypes.AccordionParams```
    - AccordionStyles ```acorntypes.AccordionStyles```
    - LabelParams ```acorntypes.LabelParams```
    - LabelStyles ```acorntypes.LabelStyles```
    - ContentParams ```acorntypes.ContentParams```
    - TimelineParams ```acorntypes.TimelineParams```
    - TestimonialParams ```acorntypes.TestimonialParams```
    - TestimonialStyles ```acorntypes.TestimonialStyles```
    - ImageParams ```acorntypes.ImageParams```
- Package **acornstyles**
    - RGX_HEXCOLOR ```Expresión regular para colores hexadecimales```
    - RGX_URL ```Expresión regular para direcciones web```
    - GetTypes ```Filled, Outlined, Pill```
    - GetWidthColumns ```Full, Quarter, Medium, ThreeQuarters, OneThird, TwoThird```
    - WithoutMargins ```Devuelve la dirección de memoria de true```
    - GetAligns ```Center, Right, Left```
    - GetColors ```Devuelve la paleta de colores de Material Design```
    - IsHexColor ```Valida que un string sea un color hexadecimal```
    - NewAcornColor ```Convierte un color hexadecimal en acorntypes.Color ```
- Package **acorn**
    - IMAGE_TIMELINE ```Imagen en base 64 de la línea de tiempo```
    - IMAGE_TESTIMONIAL ```Imagen en base 64 del icono para las recomendaciones```
    - GetBoilerplate  ```Se obtiene la plantilla general```
    - NewContent ```Se crea un contenedor```
    - NewGrid ```Se crea un grid```
    - NewRow ```Se crea un contenedor```
    - NewImage ```Se crea un componente para imagen```
    - NewSpacer ```func NewSpacer() string```
    - NewDivider ```func NewDivider(color *acorntypes.Color) string ```
    - NewAccordion ```Se crea un contenedor```
    - NewAlert ```Se crea un contenedor```
    - NewButton ```(params *acorntypes.ButtonParams)```
    - NewLabel ```Se crea un contenedor```
    - NewTestimonial ```NewTestimonial(params *acorntypes.TestimonialParams) string```
    - NewTimeline ```func NewTimeline(params []*acorntypes.TimelineParams) string```
- Package **acornmail**
    - NewAcornEmailComponents ```func NewAcornEmailComponents() *acorn.HTML```
    - MergeVariables ```func MergeVariables(temp string, variables acorntypes.AcornVariables) (string, error)```
- Archivo **LICENSE**
- Archivo **README.md**

[unreleased]: https://github.com/arskang/acornmail/tree/develop
[1.0.0-beta]: https://github.com/arskang/acornmail/releases/tag/v1.0.0-beta
