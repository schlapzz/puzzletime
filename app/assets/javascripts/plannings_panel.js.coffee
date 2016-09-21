app = window.App ||= {}
app.plannings ||= {}

app.plannings.panel = do ->
  panel = '.planning-panel'
  container = '.planning-calendar'

  setPercent = (percent) ->
    $(panel).find('#percent').val(percent)

  setDefinitive = (definitive) ->
    $(panel).find('.planning-definitive').toggleClass('active', definitive == true)
    $(panel).find('.planning-provisional').toggleClass('active', definitive == false)

    value = if definitive == null || definitive == undefined then '' else definitive.toString()
    $(panel).find('#definitive').val(value)


  definitiveChange = (event) ->
    source = $(event.target).hasClass('planning-definitive')
    current = $(panel).find('#definitive').val()
    setDefinitive(if source.toString() == current then null else source)

  position = ->
    if $(panel).length == 0 || $(panel).is(':hidden')
      return

    $(panel).position({
      my: 'right top',
      at: 'right bottom',
      of: $(container).find('.ui-selected').last(),
      within: container
    })

  cancel = ->
    app.plannings.panel.hide()
    app.plannings.selectable.clear()

  cancelOnEscape = (event) ->
    if event.key == "Escape"
      cancel()

  submit = (event) ->
    event.preventDefault()
    data = $(event.target).serializeArray()
      .reduce(((prev, curr) -> prev[curr.name] = curr.value; prev), {})
    app.plannings.service.updateSelected(data)

  deleteSelected = (event) ->
    event.preventDefault()
    # TODO: show confirmation dialog (or make it work via link_to confirm)
    app.plannings.service.deleteSelected()

  show: (selectedElements) ->
    $(panel)
      .show()
      .on('click', (event) -> event.stopPropagation())
    position()

    setPercent('')
    setDefinitive(true)

    $(panel).find('#percent').focus()

  hide: ->
    $(panel).hide()

  init: ->
    if $(panel).length == 0
      return

    $(document).on('keyup', cancelOnEscape)
    $(container).on('scroll', position)

    $(panel).find('.planning-definitive-group button').on('click', definitiveChange)
    $(panel).find('.planning-cancel').on('click', (event) ->
      $(event.target).blur()
      cancel()
    )
    $(panel).find('form').on('submit', submit)
    $(panel).find('.planning-delete').on('click', deleteSelected)

  destroy: ->
    $(document).off('keyup', cancelOnEscape)

$ ->
  app.plannings.panel.destroy()
  app.plannings.panel.init()
